fetch("/api/").catch(e => M.toast({html: e})).then(response => {
    return response.json();
}).then(data => {
    document.querySelector("#spinner").remove();
    const el = document.querySelector(".collection");
    data.forEach(c => {
        el.innerHTML += "<a class='collection-item indigo-text' href='/c/" + c + "'>" + c + "</a>";
    });
    M.AutoInit(document.body);
});
