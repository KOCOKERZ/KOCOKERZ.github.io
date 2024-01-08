import {addCSSIn} from "https://cdn.jsdelivr.net/gh/jscroot/element@0.1.5/croot.js";
import Swal from 'https://cdn.jsdelivr.net/npm/sweetalert2@11/src/sweetalert2.js';
import {id} from "../../url/config.js";


export async function main(){
    await addCSSIn("https://cdn.jsdelivr.net/npm/sweetalert2@11/dist/sweetalert2.css",id.content);
    await addCSSIn("assets/css/forum.css",id.content);
    Swal.fire({
        icon: "success",
        title: "Anda tidak memiliki pesan baru",
        showConfirmButton: false,
        timer: 1500
    });
}