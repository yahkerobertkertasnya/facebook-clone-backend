import styles from "../../src/assets/styles/navbar/navbar.module.scss";
import HomeButton from "./buttons/HomeButton.tsx";
import FriendButton from "./buttons/FriendButton.tsx";
import GroupButton from "./buttons/GroupButton.tsx";
import {useContext} from "react";
import {AuthContext} from "../context/AuthContextProvider.tsx";
import {Link} from "react-router-dom";
import {AiFillBell} from "react-icons/ai";
import {BiSolidMessageRoundedDetail} from "react-icons/bi";


export default function Navbar(){
    const auth = useContext(AuthContext);
    return (
        <>
            <div className={styles.bar}>
                <div className={styles.left}>
                    <svg xmlns="http://www.w3.org/2000/svg" width="2.5rem" height="2.5rem" fill="blue"
                         className="bi bi-facebook" viewBox="0 0 16 16">
                        <path
                            d="M16 8.049c0-4.446-3.582-8.05-8-8.05C3.58 0-.002 3.603-.002 8.05c0 4.017 2.926 7.347 6.75 7.951v-5.625h-2.03V8.05H6.75V6.275c0-2.017 1.195-3.131 3.022-3.131.876 0 1.791.157 1.791.157v1.98h-1.009c-.993 0-1.303.621-1.303 1.258v1.51h2.218l-.354 2.326H9.25V16c3.824-.604 6.75-3.934 6.75-7.951z"/>
                    </svg>
                </div>
                <div className={styles.center}>
                    <HomeButton />
                    <FriendButton />
                    <GroupButton />
                </div>
                <div className={styles.end}>
                    {/*<ProfilePicture />
                <ProfilePicture />*/}
                    <Link to={""} >
                        <div className={styles.circles}>
                            <AiFillBell
                                color={"black"}
                                size={"1.2rem"}
                            />
                        </div>
                    </Link>
                    <Link to={""} >
                        <div className={styles.circles}>
                            <BiSolidMessageRoundedDetail
                                color={"black"}
                                size={"1.2rem"}
                            />
                        </div>
                    </Link>
                    <Link to={"/user/" + auth?.username} >
                        <div className={styles.imageBox}>
                            <img src={auth?.profile ? auth?.profile : "../src/assets/default-profile.jpg"} alt={""}/>
                        </div>
                    </Link>
                </div>
            </div>
            <div className={styles.navbarSpace} >a</div>
        </>
    )
}
