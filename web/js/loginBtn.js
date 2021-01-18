$("input[type='username'],textarea").on("input",function() {
    changeColor();
});

function changeColor() {
    var filled = true;
    
    $( "input[type='username'],textarea" ).each(function() {
    if ($( this ).val() == "") {
        filled = false;
        }
    });
    if (filled) {
        $('.submit-button').css({background:'#ABFFDB'})
    } else {
        $('.submit-button').css({background:'rgb(239, 239, 239)'})
    }
}
