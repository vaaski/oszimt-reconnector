module.exports = {
  apps: [
    {
      script: "./dist/index.js",
      name: "oszimt-reconnector",
      node_args: "-r dotenv/config",
    },
  ],
}
