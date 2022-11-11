import React from "react"
import Link from "next/link"

function Header(){
    return(
        <div className="p-5 bg-stura-red">
            <Link href="/">
                <p className="font-bold">Antrag Rückerstattung anteilige Semesterbeiträge</p>    
            </Link>
        </div>
    )

}

export default Header