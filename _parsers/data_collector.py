from distutils.command.upload import upload
import os
import logging
import requests
import json
import time
from datetime import datetime, timedelta
from bs4 import BeautifulSoup
from urllib.request import Request, urlopen
from pymongo import MongoClient


db = os.environ.get("PYTHON_MONGO_URL")
db_url = f"mongodb://{db}:27017/"
client = MongoClient(db_url)
db = client.did_homa

news_collection_liga = db["news_collection_liga"]
news_collection_dou = db["news_collection_dou"]
users_collection = db['users_collection']

class NewsCollector:

    @staticmethod
    def get_liga_news() -> list:
        site = "https://news.liga.net/ua?utm_source=ua-news"
        hdr = {'User-Agent': 'Mozilla/5.0'}
        req = Request(site, headers=hdr)
        page = urlopen(req)
        soup = BeautifulSoup(page, "html.parser")
        items = soup.find_all(class_='news-nth-title')

        post_urls = []

        for i in items:
            a_tags = i.find_all('a')
            for a in a_tags:
                if "https://" in a.get('href') and len(str(a.getText())) > 30:
                    post_urls.append(a.get('href'))

        
        post_urls = post_urls[:20] 
        upload_time = int(time.time())
        for url in post_urls:

            mongo_obj = news_collection_liga.find_one({"Url": url})
            if mongo_obj is None:
                print("New Rows:", {"Url": url, "UploadTime":time.time()})
                news_collection_liga.insert_one({"Url": url, "UploadTime":upload_time})
            else:
                continue

        return post_urls

    @staticmethod
    def get_dou_news() -> list:

        site = "https://dou.ua/lenta/"
        hdr = {'User-Agent': 'Mozilla/5.0'}
        req = Request(site, headers=hdr)
        page = urlopen(req)
        soup = BeautifulSoup(page, "html.parser")
        items = soup.find_all(class_='title')

        post_names = []
        post_urls = []

        for i in items:
            a_tags = i.find_all('a')
            for a in a_tags:
                if "https://" in a.get('href'):
                    post_urls.append(a.get('href'))

        post_urls = post_urls[:20] 
        upload_time = int(time.time())
        for url in post_urls:

            mongo_obj = news_collection_dou.find_one({"Url": url})
            if mongo_obj is None:
                print("New Rows:", {"Url": url, "UploadTime":time.time()})
                news_collection_dou.insert_one({"Url": url, "UploadTime":upload_time})
            else:
                continue

        return post_urls


    @staticmethod
    def get_all_news():
        NewsCollector.get_liga_news()
        NewsCollector.get_dou_news()


    @staticmethod
    def clean_old_data():
        t = time.time() - 86400
        query = { "UploadTime": {"$lte": t} }

        for collection in [news_collection_liga, news_collection_dou]:
            collection.delete_many(query)


if __name__ == "__main__":
    t = time.time() - 86400
    myquery = { "UploadTime": {"$lte": t} }
    res = news_collection_liga.delete_many(myquery)