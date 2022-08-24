import debug from "debug"
const log = debug("oszimt-reconnector")

import got from "got"
import { load } from "cheerio"
import tough from "tough-cookie"

const jar = new tough.CookieJar()

const SPEED = 3e3
const OSZIMT_USERNAME = process.env.OSZIMT_USERNAME ?? ""
const OSZIMT_PASSWORD = process.env.OSZIMT_PASSWORD ?? ""
const OSZIMT_ADDR = "https://wlan-login.oszimt.de/logon/cgi/index.cgi"
const LOGON_BUTTON = "++Login++"

const wait = (t: number): Promise<void> => new Promise(r => setTimeout(r, t))

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

const pingLoop = async () => {
  const online = await isLoggedIn()

  if (!online) {
    await logIn()
  }

  await wait(SPEED)
  pingLoop()
}

pingLoop()
