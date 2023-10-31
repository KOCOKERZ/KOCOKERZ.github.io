import { postWithToken } from "https://jscroot.github.io/api/croot.js";
import {setInner,getValue} from "https://jscroot.github.io/element/croot.js";
import {setCookieWithExpireHour} from "https://jscroot.github.io/cookie/croot.js";

export default function PostSignUp(){
    let target_url = "https://us-central1-core-advice-401502.cloudfunctions.net/REGIS-GIS";
    let tokenkey = "token";
    let tokenvalue = "f10edefd08ff6145336a105ad8575d5beee32b93ecda886a0dd2101b50ed1c15";
    let datainjson = {
        "username": getValue("username"),
        "password": getValue("password")
    }

    postWithToken(target_url,tokenkey,tokenvalue,datainjson,responseData);

}