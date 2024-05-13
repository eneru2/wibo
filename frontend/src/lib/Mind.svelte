<script>
  let formats = {
    webp: true,
  }
  let indenting_amount = 4;
  let mime_types = false;
  let path_on_code = false;
  let absolute_path = false;
  let path_on_code_string = "tomato";
  let one_liners = false;
  let string;

  function indent(times = 1){
    let inner_spaces = ""
    for(let i = 0; i < indenting_amount*times; i++){
      inner_spaces += ` `
    }
    return inner_spaces
  }
//   let prod1 = `<img src="/${output_name}.${format}" alt="${alt}">`
  
//   let prod2 = `
// <picture>
//   <source srcset="/${output_name}.${first_format}" type="image/${first_format}">
//   <img src="/${output_name}.${second_format}" alt="${alt}">
// </picture>` 

//   let prod3 = `
// <picture>
//   <source srcset="/${output_name}.${first_format}" type="image/${first_format}">
//   <source srcset="/${output_name}.${second_format}" type="image/${second_format}">
//   <img src="/${output_name}.${third_format}" alt="${alt}">
// </picture>` 

  let happy = `<picture>
  <source srcset="/a.webp" type="image/webp">
  <source srcset="/a.webp" type="image/webp">
  <img src="/a.jpg" alt="mal">
</picture>`

  function initial_indenting(){
    const string_start = /^.+</
    const end_of_tag_regex = />\n  </g
    const end_of_tag_replace = '>\n' + indent() + '<'
    const res =
    happy
      .replace(end_of_tag_regex, end_of_tag_replace)
      .replace(string_start, "<")
    alert(res)
  }

  function append_path(){
    const src_regex = /src="\//g
    const srcset_regex = /set="\//g
    if(absolute_path){
      const src_replace = `src="/${path_on_code_string}/`
      const srcset_replace = `set="/${path_on_code_string}/`

      const res = 
      happy
        .replace(srcset_regex, srcset_replace)
        .replace(src_regex, src_replace)
        alert(res)
    } else {
      const src_replace = `src="${path_on_code_string}/`
      const srcset_replace = `set="${path_on_code_string}/`

      const res = 
      happy
        .replace(srcset_regex, srcset_replace)
        .replace(src_regex, src_replace)
        alert(res)
    }
  }

  function delete_one_liners(){
    const srcset_regex = /e s/g
    const srcset_replace = 'e\n' + indent(2) + 's'
    
    const mime_type_regex = /" t/g
    const mime_type_replace = '"\n' + indent(2) + 't'

    const src_regex = /g s/g
    const src_replace = 'g\n' + indent(2) + 's'

    const alt_regex = /" a/g
    const alt_replace = '"\n' + indent(2) + 'a'

    const result =
    happy
      .replace(srcset_regex, srcset_replace)
      .replace(mime_type_regex, mime_type_replace)
      .replace(src_regex, src_replace)
      .replace(alt_regex, alt_replace)
    alert(happy)
    alert(result)
  }

  function delete_mimes(){
    const mime_regex = / type.*2?"/g
    const result = happy.replace(mime_regex, "")
    alert(result)
  }
  function delete_absolute_paths(){

    const srcset_regex = /srcset="\//g
    const srcset_replace = 'srcset="'

    const src_regex = / src="\//g
    const src_replace = 'src="'
    const result = 
    happy
      .replace(src_regex, src_replace)
      .replace(srcset_regex, srcset_replace)
    alert(result)
  }
  // function alp(){
  //   const srcset_regex = /e s/g
  //   const srcset_replace = 'e\n' + indent(2) + 's'
    
  //   const mime_regex = /" t/
  //   const mime_newline = '"\n' + indent(2) + 't'
  //   const delete_mime_type = / type.*2?"/g

  //   const srcset_regex = /srcset="\//g
  //   const srcset_replace = 'srcset="'

  //   const src_absolute_path_regex = / src="\//g
  //   const delete_src_absolute_path = 'src="'
  //   // alert(happy)
  //   // const res = happy.replace(reg, target)
  //   // alert(res)
  //   const newq = happy.replace(delete_mime_type, "")
  //   alert(newq)
  // }
</script>
<button on:click={initial_indenting}>fur</button>