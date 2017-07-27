$.ajaxSetup({
  beforeSend: function(xhr) {
    xhr.setRequestHeader('Authenticity-Token', $('meta[name=X-CSRF-Token]').attr('content'));
  }
});

$(function() {
  $('a[data-method="delete"]').click(function(e) {
    e.preventDefault();
    var next = $(this).data('next');
    if (confirm($(this).data('confirm'))) {
      $.ajax({
        url: $(this).attr('href'),
        method: 'delete',
        success: function(rst) {
          window.location.href = next;
        },
        error: function(xhr) {
          // alert(xhr.responseText);
          console.log(xhr);
        }
      });
    }
  });
});
