// 分页插件 依赖 jquery

/**
 * size: 分页 size ，每页数据条数 默认为 10
 * total: 页码总条数
 * callback: 回调函数
 *
 * 8 页以上显示上下页，8 页以内不显示。
 *
 * */

function Pagination($element, size, total, callback) {

    // 页码宽度
    this.numbers = 8;

    // 分页容器
    this.element = $element;

    // 每页数据条数
    this.size = size || 10;

    // 数据总条数
    this.total = total;

    // 分页回调函数
    this.callback = callback || null;

    // 内部分页信息
    this.current = 1;
    this.count = total % size > 0 ? total / size + 1 : total / size;
    this.pagination = this.getPagination(this.count);

    this.validate();
}

Pagination.prototype = {
    constructor: Pagination,
    init: function () {
        // 只有一页的情况下，不显示分页
        if (this.isShowPagination()) {
            return;
        }

        var domString = this.renderPageElements();

        // 如果大于页码宽度，则显示上下页
        if (this.count > this.numbers) {
            domString = this.renderPrevElement() + domString + this.renderNextElement();
        }

        this.element.html(domString);
        this.registerEventListener();

    },
    validate: function () {
        if (typeof this.element !== 'object') {
            throw new Error("$element 请传入 jquery 对象");
        }
        if (typeof this.size !== 'number' || typeof this.total !== 'number') {
            throw new Error("size 和 total 应传入数字");
        }
    },
    getPagination: function(count) {
        var numbers = Math.min(this.numbers, count);
        var pagination = [];

        for (var i = 1; i <= numbers; i++) {
            pagination.push(i);
        }

        return pagination;

    },
    isShowPagination: function() {
        return this.total < this.size;
    },
    renderPrevElement: function() {
        return '<span class="pagination-prev">上一页</span>';
    },
    renderNextElement: function() {
        return '<span class="pagination-next">下一页</span>';
    },
    renderPageElements: function () {
        var that = this;
        var domString = '<ol class="pagination-page-wrap">';

        var nodeString = this.pagination.map(function (value) {
           return value === that.current
               ? '<li class="pagination-page actived">' + value + '</li>'
               : '<li class="pagination-page">' + value + '</li>';
        }).join('');

        domString += nodeString;

        return domString + '</ol>';
    },
    registerEventListener() {
        var that = this;

        // 上一页
        $('.pagination-prev').on('click', function () {
           if (that.current === 1) {
               alert('已经是第一页了');
               return;
           }

           that.go(that.current - 1);
        });

        // 下一页
        $('.pagination-next').on('click', function () {

            if (that.current === that.count) {
                alert('已经是最后一页了');
                return;
            }

            that.go(that.current + 1);
        });

        // 直接点击
        $('.pagination-page').on('click', function () {
            that.go(Number($(this).text()));
        });
    },
    go: function(page) {
        this.current = page;

        // 不需要改变页码
        if (this.count <= this.numbers) {
            this.init();
            this.callback(this.current, this.size);

            return;
        }

        var ignoreStart = [1, 2, 3, 4];
        var ignoreEnd = [this.count - 3, this.count - 2, this.count - 1, this.count];
        var pagination = [];

        if (ignoreStart.indexOf(page) > -1) {
            for (var i = 1; i <= 8; i++) {
                pagination.push(i);
            }
            this.pagination = pagination;

        } else if (ignoreEnd.indexOf(page) > -1) {
            for (var i = 7; i >= 0; i--) {
                pagination.push(this.count - i);
            }
            this.pagination = pagination;
        } else {
            for (var i = page - 4; i <= page + 3; i++) {
                pagination.push(i);
            }
            this.pagination = pagination;
        }

        this.init();
        this.callback(this.current, this.size);
    }
}

