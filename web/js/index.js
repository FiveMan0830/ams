window.onload = function() {
    function login() {
        const username = document.getElementById("username-field").value
        const password = document.getElementById("password-field").value
        $.ajax({
            type: "POST",
            url: "http://localhost:8080/login",
            data: {
                Username: username,
                Password: password
            },
            success: function(res) {
                console.log(res)
            }
        })
    }
    $("#login-button").click(function(){
        login()
    })
}
