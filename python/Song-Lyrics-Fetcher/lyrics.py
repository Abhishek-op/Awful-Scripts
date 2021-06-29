import os
import lyricsgenius as lg
genius = lg.Genius('Insert_Client_API_Token_Here',
                   skip_non_songs=True,
                   excluded_terms=["(Remix)", "(Live)"],
                   remove_section_headers=True)
title = input('Enter the title of the song   : ')
artist = input('Enter name of the artist/band : ')

print()
filename = '-'.join((title + ' ' + artist).split()) + '-lyrics.txt'

if os.path.exists(filename): 
    print('Lyrics of "' + title + '" already saved to the file', filename)
    exit()
else:
    file = open(filename, 'w')
    try:
        song = genius.search_song(title, artist) 
        file.write(song.lyrics)
        print('Lyrics of "' + title + '"', 'found and saved to the file',
              filename)
        file.close() 
    except:
        file.close()
        os.remove(filename)
        print('Error finding the lyrics of the song "' + title + '" by',
              artist)  
