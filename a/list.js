(() => {
    $.getJSON("/api/c/" + c, data => {
        if (!data.meta.extended)
            $('.hide-extended').hide()
        $.each(data.substitutes, (a, b) => {
            if (data.meta.extended) {
                $("tbody").append(
                    "<tr class='text-lighten-2'><td>" +
                    b.hour + "</td><td>" + b.time + "</td><td>" + b.teacher +
                    "</td><td>" + b.subject + "</td><td>" + b.room + "</td><td>" + b.type.replace("Vertretung", "Substitute") +
                    "</td><td>" + b.reason + "</td></tr>");
            } else {
                $("tbody").append(
                    "<tr class='text-lighten-2'><td>" +
                    b.hour + "</td><td>" + b.teacher +
                    "</td><td>" + b.subject + "</td><td>" + b.room + "</td><td>" + b.type.replace("Vertretung", "Substitute") +
                    "</td></tr>");
            }
        });
        $("h4").html(data.meta.date.replace("Vertretungen", "Substitutes").split("/")[0]);
        $("#title").html(data.meta.class);
    }).catch(m => Materialize.toast(m.status + ": " + m.responseJSON.message));
})();