// function getFormJson(frm) {
// 	var o = {}
// 	var a = $(frm).serializeArray()
// 	$.each(a, function(){
// 		o[this.name] = this.value || '';
// 	})
// 	return o;
// }
// // 表单数据提交给服务器，服务器返回后，执行fn函数
// function ajaxSubmit(frm, fn) {
// 	var dat = getFormJson(frm)
// 	return $.ajax({
// 		url : frm.action,
// 		type: frm.method,
// 		dta : dat,
// 		success: fn
// 	})
// }


function getFormJson(frm) {
	var o = {};
	var a = $(frm).serializeArray()
	$.each(a, function(){
		o[this.name] = this.value || '';
	})
	return o;
}

function ajaxSubmit(frm, fn){
	var dat = getFormJson(frm)
	return $.ajax({
		url : frm.action,
		type: frm.method,
		data: dat,
		success: fn
	})
}