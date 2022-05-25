# telebot (Ukrainian language)
My first telegram bot on Golang. 
Funtionality:
1. Parse and show news
2. Connected to predition data
3. Connected to aphorisms data

My attempts to study GO language with developing Telegram Bot.
Bot consists of 3 componnents in docker containers:
  1. Core bot logic (Golang)
  2. News parsers (Python)
  3. Database (Mongo)

### Configuration:
Before start you need provide 2 environment files:
1. `.env` in root folder with next variables:  `TELEGRAM_APITOKEN, MONGO_URL, AWS_S3_BUCKET_NAME, AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY`
2. `.env` in parsers folder with `PYTHON_MONGO_URL` variable


### Start project:
`docker-compose build`
`docker-compose up`

### Additional info
Be aware that this project use data from AWS S3 bucket that I don't share, to establish proper work you need to create and configure your own data sources. 


### Elample
![alt text](https://github.com/AntonAks/telebot/blob/main/telebot_pic.jpg?raw=true)