
$(document).ready(function () {

    // 右侧点赞收藏
    $(".article-reletive-detial-icon").each(function (index, element) {
        // 主播打赏额外处理
        if (index === 2) {
            $(element).hover(function () {
                $(element).css({backgroundColor: '#ee5044', color: '#ffffff', borderColor: '#ee5044'}).
                    find('i').hide().next('span').show().next('img').show();
            }, function () {
                $(element).css({backgroundColor: '#ffffff', color: '#999999', borderColor: '#999999'}).
                    find('img').hide().prev('span').hide().prev('i').show();
            });
        } else {
            $(element).hover(function () {
                $(element).css({color: '#ffffff', backgroundColor: '#999999'}).find('i').hide().next('span').show();
            }, function () {
                $(element).css({color: '#999999', backgroundColor: '#ffffff'}).find('span').hide().prev('i').show();
            });
        }
    });
});