
function updateUserStateInfo() {
    $(".article-relative-detail-icon").each(function(index, element) {
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
            var color = $(element).hasClass("selected") ? '#ee5044' : '#999999';
            $(element).hover(function () {
                $(element).css({color: '#ffffff', backgroundColor: color}).find('i').hide().next('span').show();
            }, function () {
                $(element).css({color: color, backgroundColor: '#ffffff'}).find('span').hide().prev('i').show();
            });
        }
    });
}

$(document).ready(function () {

    // 右侧点赞收藏
    updateUserStateInfo();

    // 相关文章
    $(".article-relative-articles").find(".icon-angle-right").click(function () {
        var $element = $(".article-relative-article-panel-inner");
        var left = parseInt($element.css('left'), 10);

        if (left > -1200) {
            $element.animate({ left: (left - 300) + 'px' }, 300);
        }
    });

    $(".article-relative-articles").find(".icon-angle-left").click(function () {
        var $element = $(".article-relative-article-panel-inner");
        var left = parseInt($element.css('left'), 10);

        if (left < 0) {
            $element.animate({ left: (left + 300) + 'px' }, 300);
        }
    });

    // 文章点赞
    $('.article-support').on('click', function() {
       var isSelected = $(this).hasClass('selected');
       var articleId = $(this).attr('data-articleId');

       var params = { isSupported: !isSelected, articleId: articleId };

       var $that = $(this);

       var number = +($that.next('.article-relative-detail-text').text());

       $.ajax({
           type: 'POST',
           url: '/article/support',
           data: params,
           headers: {
             token: window.localStorage.getItem("token")
           },
           success: function(result) {
               if (result.rc === "0") {
                   if (!isSelected) {
                       $that.addClass("selected").next('.article-relative-detail-text').addClass("selected");
                       $that.next('.article-relative-detail-text').text(number + 1);
                   } else {
                       $that.removeClass("selected").next('.article-relative-detail-text').removeClass("selected");
                       $that.next('.article-relative-detail-text').text(number - 1);
                   }
                   updateUserStateInfo();
               } else {
                   alert(result.msg);
               }
           },
           error: function(error) {
               console.log(error);
           }
       })
    });

    // 文章收藏
    $('.article-collection').on('click', function() {
        var isSelected = $(this).hasClass('selected');
        var articleId = $(this).attr('data-articleId');

        var params = { isCollected: !isSelected, articleId: articleId };

        var $that = $(this);

        var number = +($that.next('.article-relative-detail-text').text());

        $.ajax({
            type: 'POST',
            url: '/article/collect',
            data: params,
            headers: {
                token: window.localStorage.getItem("token")
            },
            success: function(result) {
                if (result.rc === "0") {
                    if (!isSelected) {
                        $that.addClass("selected").next('.article-relative-detail-text').addClass("selected");
                        $that.next('.article-relative-detail-text').text(number + 1);
                    } else {
                        $that.removeClass("selected").next('.article-relative-detail-text').removeClass("selected");
                        $that.next('.article-relative-detail-text').text(number - 1);
                    }
                    updateUserStateInfo();
                } else {
                    alert(result.msg);
                }
            },
            error: function(error) {
                console.log(error);
            }
        })
    })

    // 关注用户
    $('.author-opecation').on('click', function() {
       var isFollowing = $(this).find('span').text() === '取消关注';
       var targetUserId = $(this).attr('data-anchor');

       var params = { isFollowing: !isFollowing, anchorId: targetUserId };

       var $that = $(this);
       var followers = +($('.author-detail').find('p').eq(1).find('span').text());

       $.ajax({
           type: 'POST',
           url: '/following/anchor',
           data: params,
           headers: {
               token: window.localStorage.getItem("token")
           },
           success: function(result) {
               if (result.rc === "0") {
                   var text = isFollowing ? '关注' : '取消关注';
                   var number = isFollowing ? followers - 1 : followers + 1;

                   $that.find('span').text(text);
                   $('.author-detail').find('p').eq(1).find('span').text(number)
               } else {
                   alert(result.msg);
               }
           },
           error: function(error) {
               console.log(error);
           }
       });
    });
});