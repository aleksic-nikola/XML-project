function crosstest() {


    $.ajax({
        type:'GET',
        crossDomain: true,
        url: 'http://localhost:5050/getmsgs',
        //headers: { 'Content-Type': 'application/json' },
        contentType : 'application/json',
        success : function(data) {
            console.log(data)
        },
        error : function(xhr, status, data) {
            console.log(xhr)
            console.log('Cant get msgs');
        }
    })
}