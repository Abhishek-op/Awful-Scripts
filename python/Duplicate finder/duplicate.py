import os

def duplicate(folder1, folder2):
    f1content = os.listdir(folder1)
    f2content = os.listdir(folder2)
    for item in f1content:
        for check in f2content:
            if item == check:
                print(item)
                
                
def find():
	d1 = input("Enter path of First Folder: ")
	d2 = input("Enter path of Second Folder: ")
	duplicate(d1,d2)
    
    
if __name__ == '__main__':
	find()