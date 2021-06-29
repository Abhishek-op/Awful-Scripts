from PyPDF2 import PdfFileWriter, PdfFileReader

print("[+] The file should be in same FOLDER as this script\n")

pdf_File_Name = input("Enter EXACT name of the PDF in this FOLDER: ")
pdf_File = pdf_File_Name + ".pdf"

pdf = PdfFileReader(pdf_File)

write_Obj = PdfFileWriter()

for i in range(pdf.getNumPages()):
    page = pdf.getPage(i)
    write_Obj.addPage(page)

owner_Password = input("Enter Password for OWNER: ")
user_Password = input("Enter Password for USER: ")
write_Obj.encrypt(user_pwd=user_Password,
                  owner_pwd=owner_Password,
                  use_128bit=True)
new_PDF_Name_Input = input("Enter new ENCRYPTED PDF name: ")
new_PDF_Name = new_PDF_Name_Input + '.pdf'
encrypted_PDF = open(new_PDF_Name, 'wb')
write_Obj.write(encrypted_PDF)
