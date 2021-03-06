let currentClass = document.location.href.substr(document.location.href.lastIndexOf("/") + 1);


fetch("/api/c/" + currentClass).catch(e => M.toast({html: e})).then(res => {
    return res.json();
}).then(data => {
    document.querySelector("#spinner").remove();
    document.querySelector("h4").innerText = data.meta.date.replace("Vertretungen", "Substitutes").split("/")[0];
    document.querySelector("#title").innerHTML = data.meta.class + ", " + data.meta.date.split("/")[1] + " - " + data.meta.updated;
    if (!data.meta.extended)
        Array.from(document.getElementsByClassName("hide-extended")).forEach(a => a.remove());
    data.substitutes.forEach(substitute => {
        // TODO: Smart fill for this
        if (data.meta.extended) {
            document.querySelector("tbody").innerHTML += "<tr class='text-lighten-2'><td>" +
                substitute.hour + "</td><td>" + substitute.time + "</td><td>" + substitute.teacher.replace("?", " => ") +
                "</td><td>" + substitute.subject.replace("?", " => ") + "</td><td>" + substitute.room + "</td><td>" + substitute.type.replace("Vertretung", "Substitute") +
                "</td><td>" + substitute.notes + "</td><td>" + substitute.reason + "</td></tr>";
        } else {
            document.querySelector("tbody").innerHTML += "<tr class='text-lighten-2'><td>" +
                substitute.hour + "</td><td>" + substitute.teacher.replace("?", " => ") +
                "</td><td>" + substitute.subject.replace("?", " => ") + "</td><td>" + substitute.room + "</td><td>" + substitute.type.replace("Vertretung", "Substitute") +
                "</td><td>" + substitute.notes + "</td></tr>";
        }
    });
});
