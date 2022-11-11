import React from "react"
import Link from "next/link"

function Header(){
    return(
        <div className="p-5 bg-stura-yellow">
            <Link href="/">
                <p className="font-bold">Admin Panel</p>    
            </Link>
        </div>
    )

}

export default Header