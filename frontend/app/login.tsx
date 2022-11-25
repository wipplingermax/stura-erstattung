'use client'

import { useState } from "react";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup"
import { useRouter } from "next/navigation";

const schema = yup.object().shape({
    uniID: yup.string().matches(/^[a-zA-Z]{2}[0-9]{3}$/,"Die Uni-ID hat das Format 'jb007'").required("Uni-ID benötigt"),
    password: yup.string().required("Passwort benötigt"),
})

export default function Login() {

    const {register, handleSubmit, formState: {errors}} = useForm({resolver: yupResolver(schema)});  
    const router = useRouter();
    const [loginError, setLoginError] = useState(false);

    function submitRequest(data: any){

        if (auth(data)){      
            router.push("/request");
        } else {
            console.log("login error")
        }
    }

    function auth(data :any){
        if (!data) {
            return true
        } else {
            setLoginError(true);
            return false
        }
     }

    return (
        <div className="bg-white rounded shadow-xl p-8 mt-6 mb-6 w-96">   
            <div className="font-bold text-center text-xl">Anfrage Erstattung</div>
            <div className="font-bold text-center text-xl">Semesterbeitrag</div>  
            
            { loginError == false ? 
                (
                    <form onSubmit={handleSubmit(submitRequest)} className="pt-8">
                        <div className="grid gap-4 grid-cols-1">
                            <div>
                                <input type="text" placeholder="Uni-ID" className="h-10 border mt-1 rounded px-4 w-full bg-gray-50" {...register("uniID")} aria-invalid={errors.uniID ? "true" : "false"}/>
                                {errors.uniID? <p className="text-stura-red text-xs p-2">{errors.uniID!.message?.toString()}</p>:""}
                            </div>
                            <div>
                                <input type="password" placeholder="Passwort" className="h-10 border mt-1 rounded px-4 w-full bg-gray-50" {...register("password")} aria-invalid={errors.password ? "true" : "false"}/>
                                {errors.password? <p className="text-stura-red text-xs p-2">{errors.password!.message?.toString()}</p>:""}
                            </div>
                        </div>
                        <div className="flex justify-center">
                            <input type="submit" name="submit" value="Login" className="bg-stura-red hover:bg-stura-red-hover text-white font-bold py-2 px-4 rounded mt-10"/>
                        </div>
                    </form>
                ) : (
                    <div className="flex flex-col justify-center">
                        <div className="py-8 flex justify-center">Login fehlgeschlagen</div>
                        <div className="flex justify-center">
                            <button className="bg-stura-red hover:bg-stura-red-hover text-white font-bold py-2 px-4 rounded"
                            onClick={() => setLoginError(false)}>Erneut versuchen</button>
                        </div>
                    </div>
                ) 
            }
        </div>
    );

  }