import kivy
from kivy.app import App
from kivy.uix.boxlayout import BoxLayout
from kivy.lang import Builder
from PIL import Image

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
        img_path = f
        d = str(f).replace('.jpg', '.pdf')
        
        image1 = Image.open(f)
        im1 = image1.convert('RGB')
        im1.save(d)
        
  
        
        
        


class MyApp(App):
    def build(self):
        return MyWidget()

if __name__ == '__main__':
    MyApp().run()