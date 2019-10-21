use std::collections::HashMap;
use std::ffi::CString;
use std::os::raw::c_char;

#[link(name = "curl")]
#[link(name = "fetch")]
extern {
  fn fetch_init();
  fn fetch_cleanup();
  fn fetch(url: *const c_char, filename: *const c_char) -> u64;
}

fn main() {
  let urls: HashMap<&str, &str> = [
    ("https://apple.com/favicon.ico", "apple.ico"),
    ("https://google.com/favicon.ico", "google.ico"),
  ].iter().cloned().collect();

  unsafe { fetch_init() };

  for (&url, &file) in &urls {
    let c_url = CString::new(url).expect("failed to copy url");
    let c_file = CString::new(file).expect("failed to copy file");

    unsafe { fetch(c_url.as_ptr(), c_file.as_ptr()) };
  }

  unsafe { fetch_cleanup() };
}
