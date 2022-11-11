import React from "react"
import Link from "next/link"

function Header(){
    return(
        <div className="w-full p-5 text-white shadow-lg bg-white fixed">
            <Link href="/">
                <img src="/StuRa_Logo.svg" alt="Logo" className="h-10" />    
            </Link>
        </div>
    )

}

export default Header