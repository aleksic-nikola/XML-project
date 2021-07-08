function search() {

    var searchinput = $("#search_input").val()

    localStorage.setItem("query", searchinput)

    window.location.href = "search.html"
}