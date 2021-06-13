
import kivy
import os
import zipfile
from kivy.app import App
from kivy.uix.boxlayout import BoxLayout
from kivy.lang import Builder

import os

Builder.load_string("""
<MyWidget>:
    id: my_widget
    FileChooserIconView:
        id: filechooser
        on_selection: my_widget.selected(filechooser.selection)
""")

class MyWidget(BoxLayout):
    def open(self, path, filename):
        with open(os.path.join(path, filename[0])) as f:
            print(f.read())

    def selected(self, filename):
        s = str(filename[0]).replace('[', '')
        f = str(s).replace(']', '')
        d = str(f).replace('.zip', '')
        os.mkdir(d)
        print("selected:"+f)
        with zipfile.ZipFile(f, 'r') as zip_ref:
        	zip_ref.extractall(d)
        
class MyApp(App):
    def build(self):
        return MyWidget()

if __name__ == '__main__':
    MyApp().run()