import React from "react"
import Link from "next/link"

function Header(){
    return(
        <div className="w-full p-4 text-white shadow-lg bg-white fixed">
            <img src="/StuRa_Logo.svg" alt="Logo" className="h-10" />    
        </div>
    )

}

export default Header