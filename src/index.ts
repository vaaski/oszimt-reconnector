import debug from "debug"
const log = debug("oszimt-reconnector")

import got from "got"
import { load } from "cheerio"
import tough from "tough-cookie"
import notifier from "node-notifier"
import wifi from "node-wifi"
import anybar from "anybar"

import { dirname, join } from "path"
import { fileURLToPath } from "url"
import { readFileSync } from "fs"

const __dirname = dirname(fileURLToPath(import.meta.url))

const packageJson = readFileSync(join(__dirname, "..", "package.json"))
const version = JSON.parse(packageJson.toString()).version as string

const SPEED = 3e3
const OSZIMT_USERNAME = process.env.OSZIMT_USERNAME ?? ""
const OSZIMT_PASSWORD = process.env.OSZIMT_PASSWORD ?? ""
const OSZIMT_ADDR = "https://wlan-login.oszimt.de/logon/cgi/index.cgi"
const LOGON_BUTTON = "++Login++"
const COMPATIBLE_NETWORKS = ["OSZIMTSchueler", "OSZIMTBesucher"]

const jar = new tough.CookieJar()

const wait = (t: number): Promise<void> => new Promise(r => setTimeout(r, t))
const notify = (text: string, timeout = 5) => {
  try {
    return notifier.notify({
      title: `oszimt-reconnector v${version}`,
      message: text,
      open: OSZIMT_ADDR,
      timeout,
    })
  } catch (error) {
    log(JSON.stringify(error))
  }
}

const isLoggedIn = async (): Promise<boolean> => {
  const response = await got(OSZIMT_ADDR)
  const $ = load(response.body)
  const loggedIn = !!$(".logged-in").length

  log(`is logged in: ${loggedIn}`)
  return loggedIn
}

const getLoginToken = async () => {
  const loggedIn = await isLoggedIn()
  if (loggedIn) {
    log("already logged in")
    return ""
  }

  const response = await got(OSZIMT_ADDR)
  const $ = load(response.body)

  const token = $(`input[name=ta_id]`).attr("value")
  if (!token) throw "couldn't get token"

  const cookies = response.headers["set-cookie"]
  if (cookies) {
    for (const cookie of cookies) {
      await jar.setCookie(cookie, OSZIMT_ADDR)
    }
  }

  log(`login token: ${token}`)
  return token
}

const logIn = async () => {
  const loggedIn = await isLoggedIn()
  if (loggedIn) return log("already logged in")
  else log("logging in")

  await got.post({
    url: OSZIMT_ADDR,
    form: {
      ta_id: await getLoginToken(),
      uid: OSZIMT_USERNAME,
      pwd: OSZIMT_PASSWORD,
      voucher_logon_btn: LOGON_BUTTON,
    },
    headers: {
      Cookie: await jar.getCookieString(OSZIMT_ADDR),
    },
  })

  log("after login")
}

let lastWasCorrect = true
const pingLoop = async (): Promise<void> => {
  anybar("filled")

  try {
    const correctNetwork = await isCorrectNetwork()
    if (!correctNetwork) {
      log("not on correct network")
      if (lastWasCorrect) {
        anybar("exclamation")
        notify("not on correct network")
      }
      lastWasCorrect = false

      await wait(SPEED)
      return pingLoop()
    }

    if (!lastWasCorrect) notify("back on correct network")
    lastWasCorrect = true

    const online = await isLoggedIn()

    if (!online) {
      anybar("green")
      notify("logging in...", 10)
      await logIn()
      notify("logged in", 2)
    }
  } catch (error) {
    log(error)
    notify(JSON.stringify(error), 10)
  }

  anybar("hollow")
  await wait(SPEED)
  pingLoop()
}

const isCorrectNetwork = async () => {
  const connections = await wifi.getCurrentConnections()
  for (const connection of connections) {
    if (COMPATIBLE_NETWORKS.includes(connection.ssid)) {
      return true
    }
  }

  return false
}

wifi.init({ iface: null })

isCorrectNetwork().then(correct => {
  if (correct) notify("starting ping-loop", 2)
})

pingLoop()

let exiting = false
const onQuit = async () => {
  if (exiting) return
  exiting = true

  await anybar("question")
  process.exit()
}

process.on("exit", onQuit)
process.on("SIGINT", onQuit)
