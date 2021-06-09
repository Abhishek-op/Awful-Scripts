
function bin2dec() {
    let binaryNum = document.getElementsByClassName("number1")[0].value;
    var binary = parseInt(binaryNum,2)
    
    document.getElementsByClassName("result")[0].value = binary;
}
function hex2dec() {
    let binaryNum = document.getElementsByClassName("number1")[0].value;
    var binary = parseInt(binaryNum,16)
    
    document.getElementsByClassName("result")[0].value = binary;
}
function oct2dec() {
    let binaryNum = document.getElementsByClassName("number1")[0].value;
    var binary = parseInt(binaryNum,8)
    
    document.getElementsByClassName("result")[0].value = binary;
}
        
        
        
     
    