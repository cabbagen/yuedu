
$(document).ready(function () {
    var pagination = new Pagination($('.channel-pagination'), 10, 120, function (current, size) {
        console.log(current, size);
    });

    pagination.init();
});