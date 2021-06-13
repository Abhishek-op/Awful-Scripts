import bs4
from bs4 import BeautifulSoup as soup
from urllib.request import urlopen
def newss():
    
    news_url="https://news.google.com/news/rss"
    Client=urlopen(news_url)
    xml_page=Client.read()
    Client.close()
    soup_page=soup(xml_page,"html.parser")
    news_list=soup_page.findAll("item")
# Print news title, url and publish date
    for news in news_list:
    	print(news.title.text)
    	print(news.link.text)
newss()