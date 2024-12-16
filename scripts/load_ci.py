from pymongo import MongoClient
import json

client = MongoClient('mongodb://root:example@mongodb:27017/mydb?authSource=admin')

database   = client['serenesong']
collection = database['Ci']
collection.delete_many({})

for iter in range(0, 22):
    with open(f'./database/songci/ci.song.{iter*1000}.json', 'r', encoding='utf-8') as f:
        data = json.load(f)
        
    new_data = []
    for item in data:
        new_item = {}
        for key in item:
            if   key == 'paragraphs':
                new_content = []
                for sentence in item[key]:
                    if sentence == " >> " or sentence == "词牌介绍":
                        continue
                    new_content.append(sentence)
                new_item['content'] = new_content
            elif key == 'rhythmic':
                if '・' in item[key]:
                    cipai_list = []
                    for cipai in item[key].split('・'):
                        cipai_list.append(cipai)
                    new_item['cipai'] = cipai_list
                else:
                    new_item['cipai'] = [item[key]]
            elif key == 'prologue':
                new_item['xiaoxu'] = item[key]
            else:
                new_item[key] = item[key]
        new_data.append(new_item)
        
    # client = MongoClient('mongodb://root:example@mongodb:27017/mydb?authSource=admin')
    # database   = client['serenesong']
    # collection = database['Ci']
    # collection.delete_many({})
    result = collection.insert_many(new_data)

    print("Data inserted successfully.")
    # client.close()
    
with open(f'./database/songci/ci.song.2019y.json', 'r', encoding='utf-8') as f:
    data = json.load(f)
    
new_data = []
for item in data:
    new_item = {}
    for key in item:
        if key == 'paragraphs':
            new_item['content'] = item[key]
        elif key == 'rhythmic':
            if '・' in item[key]:
                cipai_list = []
                for cipai in item[key].split('・'):
                    cipai_list.append(cipai)
                new_item['cipai'] = cipai_list
            else:
                new_item['cipai'] = [item[key]]
        elif key == 'prologue':
            new_item['xiaoxu'] = item[key]
        else:
            new_item[key] = item[key]
    new_data.append(new_item)
    
# client = MongoClient('mongodb://root:example@mongodb:27017/mydb?authSource=admin')
# database   = client['admin']
# collection = database['Ci']
# collection.delete_many({})
result = collection.insert_many(new_data)

print("Data inserted successfully.")
client.close()