'use client'

export default function RequestForm() {
    return (
       <div className="form-wrapper bg-white rounded shadow-lg p-4 px-4 md:p-8 mb-6">
            <form className="grid gap-4 grid-cols-1 lg:grid-cols-2" >
                <label className="font-bold lg:col-span-2">Personenbezogene Daten</label>
                <input type="text" name="first_name" required placeholder="Vorname" className="h-10 border mt-1 rounded px-4 w-full bg-gray-50"/>
                <input type="text" name="last_name" required placeholder="Nachname" className="h-10 border mt-1 rounded px-4 w-full bg-gray-50"/>
                <input type="text" name="matriculation_number" required placeholder="Matrikelnummer" className="h-10 border mt-1 rounded px-4 w-full bg-gray-50"/>
                <input type="text" name="uni-id" required placeholder="Uni-ID" className="h-10 border mt-1 rounded px-4 w-full bg-gray-50"/>
                <input type="text" name="e-mail" required placeholder="E-Mail" className="lg:col-span-2 h-10 border mt-1 rounded px-4 w-full bg-gray-50"/>
                <input type="text" name="phone" placeholder="optional: Telefonnummer" className="lg:col-span-2 h-10 border mt-1 rounded px-4 w-full bg-gray-50"/>
                
                <label className="font-bold lg:col-span-2 pt-6 ">Bankverbindng</label>
                <input type="text" name="iban" required placeholder="IBAN" className="lg:col-span-2 h-10 border mt-1 rounded px-4 w-full bg-gray-50"/>
                <input type="text" name="bic" placeholder="optional: BIC" className="lg:col-span-2 h-10 border mt-1 rounded px-4 w-full bg-gray-50"/>
                <input type="text" name="account_owner" placeholder="Kontoinhaber" className="lg:col-span-2 h-10 border mt-1 rounded px-4 w-full bg-gray-50"/>
            </form>
            <div className="text-center mt-5 lg:text-end">
                <button className="bg-stura-red hover:bg-stura-red-hover text-white font-bold py-2 px-4 rounded">Absenden</button>
            </div>
            
       </div>
    );
  }