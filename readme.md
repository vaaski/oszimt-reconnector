<h1 align="center">osz-imt reconnector</h1>

#### install
```bash
git clone https://github.com/vaaski/oszimt-reconnector
cd oszimt-reconnector
npm ci
```

----

#### configure credentials

- copy .env.example to .env
- fill out username and password

----

#### run once
```bash
node dist
```

#### run forever
```bash
npm i -g pm2
pm2 startup
pm2 start pm2.config.json
pm2 save
```