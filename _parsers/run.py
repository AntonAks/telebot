import logging
import asyncio
import aioschedule
from data_collector import NewsCollector
from random import choice


async def get_news():
    NewsCollector.get_all_news()

async def clean_db():
    NewsCollector.clean_old_data()


async def scheduler():
    aioschedule.every(15).minutes.do(get_news)
    aioschedule.every().day.at("03:00").do(clean_db)

    while True:
        await aioschedule.run_pending()
        await asyncio.sleep(0.5)


if __name__ == '__main__':

    # INITIAL POPULATE
    NewsCollector.get_all_news()

    loop = asyncio.get_event_loop()
    # SCHEDULER
    loop.run_until_complete(scheduler())