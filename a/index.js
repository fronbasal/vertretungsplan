// IndexJS
(() => {
    const e = $('select');
    $.getJSON("/api", data => {
        $.each(data, (a, b) => {
            $("#today").append("<option value='" + b + "'>" + b + "</option>");
        });
        e.material_select();
        e.change(function () {
            window.location.pathname = "/c/" + $(this).val();
        });
    }).catch(e => console.log(e));
})();
