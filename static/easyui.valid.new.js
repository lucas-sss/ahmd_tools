(function ($) {
    /**
     * Copyright (c) 2009-2015 RCM
     *
     * 新增校验规则
     *
     */
    $.extend($.fn.validatebox.defaults.rules, {
        checkNumAndChar: {
            validator: function (value, param) {
                return /^\w+$/.test(value);
            },
            message: '输入为数字、字母和下划线组合'
        }
    });
})(jQuery);



