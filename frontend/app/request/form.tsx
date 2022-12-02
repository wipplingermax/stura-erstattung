'use client'

import React from "react";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import {useMutation} from "react-query"
import * as yup from "yup"
import { electronicFormatIBAN, isValidIBAN, isValidBIC } from 'ibantools'

const api = "http://127.0.0.1:8080/v1/request"

const schema = yup.object().shape({
    firstName: yup.string().required("Vorname benötigt"), 
    lastName: yup.string().required("Nachname benötigt"),
    matriculationNumber: yup.string().matches(/^[0-9]{7}$/,"Die Matrikelnummer hat das Format '1234567'").required("Matrikelnummer benötigt"),
    uniID: yup.string().matches(/^[a-zA-Z]{2}[0-9]{3}$/,"Die Uni-ID hat das Format 'jb007'").required("Uni-ID benötigt"),
    email: yup.string().email("E-Mail inkorrekt").required("E-Mail wird benötigt"),
    phone: yup.string().matches(/^[0-9\+\-\(\) ]{0,20}$/,"Invalide Telefonnummer"),
    iban: yup.string().required("IBAN benötigt").test('iban','IBAN ist inkorrekt', value => {return isValidIBAN(electronicFormatIBAN(value!)!)}),
    bic: yup.string().test('bic','BIC ist inkorrekt', value => {return value == "" || isValidBIC(value!)}),
    accountOwner: yup.string().required("Kontoinhaber benötigt").max(40,"Kontoinhaber darf maximal 40 Zeichen lang sein"),
    verficationNumber: yup.string().matches(/^[a-zA-Z]{12}$/, "Die Verfikationsnummer hat das Format'ABCD EFGH IJKL'"),
})

export default function RequestForm() {
    const {register, handleSubmit, formState: {errors}} = useForm({resolver: yupResolver(schema)});


    function submitRequest(data: any){
        console.log(data);
    }
   
    return (
            <form onSubmit={handleSubmit(submitRequest)} className="bg-white rounded shadow-xl p-4 px-4 md:p-8 mb-6">
                <div className="grid gap-4 grid-cols-1 lg:grid-cols-2">
                    <label className="font-bold lg:col-span-2">Personenbezogene Daten</label>
                    <div>
                        <input type="text" placeholder="Vorname" className="h-10 border mt-1 rounded px-4 w-full bg-gray-50" {...register("firstName")} />
                        {errors.firstName? <p className="text-stura-red text-xs p-2">{errors.firstName.message?.toString()}</p>:""}
                    </div>
                    <div>
                        <input type="text" placeholder="Nachname" className="h-10 border mt-1 rounded px-4 w-full bg-gray-50" {...register("lastName")}/>
                        {errors.lastName? <p className="text-stura-red text-xs p-2">{errors.lastName.message?.toString()}</p>:""}
                    </div>
                    <div>
                        <input type="text" placeholder="Matrikelnummer" className="h-10 border mt-1 rounded px-4 w-full bg-gray-50" {...register("matriculationNumber")}/>
                        {errors.matriculationNumber? <p className="text-stura-red text-xs p-2">{errors.matriculationNumber.message?.toString()}</p>:""}
                    </div>
                    <div>
                        <input type="text" placeholder="Uni-ID" className="h-10 border mt-1 rounded px-4 w-full bg-gray-50" {...register("uniID")}/>
                        {errors.uniID? <p className="text-stura-red text-xs p-2">{errors.uniID.message?.toString()}</p>:""}
                    </div>
                    
                    <input type="text" placeholder="E-Mail" className="lg:col-span-2 h-10 border mt-1 rounded px-4 w-full bg-gray-50" {...register("email")}/>
                    {errors.email? <p className="text-stura-red text-xs p-2">{errors.email.message?.toString()}</p>:""}
                    
                    <input type="text" placeholder="optional: Telefonnummer" className="lg:col-span-2 h-10 border mt-1 rounded px-4 w-full bg-gray-50" {...register("phone")}/>
                    {errors.phone? <p className="text-stura-red text-xs p-2">{errors.phone.message?.toString()}</p>:""}
                    
                    <label className="font-bold lg:col-span-2 pt-6 ">Bankverbindung</label>
                   
                    <input type="text" placeholder="IBAN" className="lg:col-span-2 h-10 border mt-1 rounded px-4 w-full bg-gray-50" {...register("iban")}/>
                    {errors.iban? <p className="text-stura-red text-xs p-2">{errors.iban.message?.toString()}</p>:""}                  
                    
                    <input type="text" placeholder="optional: BIC" className="lg:col-span-2 h-10 border mt-1 rounded px-4 w-full bg-gray-50" {...register("bic")}/>
                    {errors.bic? <p className="text-stura-red text-xs p-2">{errors.bic.message?.toString()}</p>:""}
                    
                    <input type="text" placeholder="Kontoinhaber" className="lg:col-span-2 h-10 border mt-1 rounded px-4 w-full bg-gray-50" {...register("accountOwner")}/>
                    {errors.accountOwner? <p className="text-stura-red text-xs p-2">{errors.accountOwner.message?.toString()}</p>:""}

                    <label className="font-bold lg:col-span-2 pt-6 ">Legitimationsnachweis</label>

                    <div>
                        <input type="text" placeholder="Verfikationsnummer" className="h-10 border mt-1 rounded px-4 w-full bg-gray-50" {...register("verficationNumber")} />
                        {errors.verficationNumber? <p className="text-stura-red text-xs p-2">{errors.verficationNumber.message?.toString()}</p>:""}
                    </div>
                    <div>
                        <p className="text-xs px-4 mt-1 h-10">Wo finde ich meine Verfikationsnummer?</p>
                    </div>

                    
                </div>
                <div className="flex justify-center lg:justify-end">
                    <input type="submit" name="submit" value="Absenden" className="bg-stura-red hover:bg-stura-red-hover text-white font-bold py-2 px-4 rounded mt-10"/>
                </div>
            </form>
    );
  }