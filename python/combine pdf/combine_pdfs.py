import PyPDF2, os

pdf_files=[filename for filename in os.listdir('.') if filename.endswith('.pdf')]
pdf_files.sort()
count_merged_files = 0
pdfWriter = PyPDF2.PdfFileWriter()
for filename in pdf_files:
	print(f'Merging \'{filename}\' to main pdf file.')
	count_merged_files+=1
	
	pdfFileObj = open(filename,'rb')
	pdfReader = PyPDF2.PdfFileReader(pdfFileObj,strict=False)
	for page_num in range(0,pdfReader.numPages):
		pageObj = pdfReader.getPage(page_num)
		pdfWriter.addPage(pageObj)
pdfOutput = open('output.pdf','wb')
pdfWriter.write(pdfOutput)

pdfOutput.close()
print('')
print(f'[+] {count_merged_files} files combined.')
