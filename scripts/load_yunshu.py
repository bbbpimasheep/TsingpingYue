from pymongo import MongoClient
import json

with open('./database/rhymes/Pingshui_Rhyme.json', 'r', encoding='utf-8') as f:
    data = json.load(f)
    
new_data = {}
for tone, info in data.items():
    new_data[tone] = info
    
client = MongoClient('mongodb://root:example@mongodb:27017/mydb?authSource=admin')

database   = client['serenesong']
collection = database['PingshuiYun']
collection.delete_many({})
result     = collection.insert_one(new_data)

print("Data inserted successfully.")
client.close()