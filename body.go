package main

// ObtenerClienteBody used on the show client request
const ObtenerClienteBody = `<?xml version=\"1.0\" encoding=\"UTF-8\"?> 
<SOAP-ENV:Envelope
    xmlns:SOAP-ENV=\"http://schemas.xmlsoap.org/soap/envelope/\"
    xmlns:ns1=\"https://plazavip.clarodrive.com/dla/soap/\"
    xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\"
    xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"
    xmlns:SOAP-ENC=\"http://schemas.xmlsoap.org/soap/encoding/\" SOAP-ENV:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">
    <SOAP-ENV:Body>
        <ns1:obtenerCliente>
            <id xsi:type=\"xsd:int\">%s</id>
        </ns1:obtenerCliente>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

// ActualizarClienteBody used on the update client request
const ActualizarClienteBody = `<?xml version=\"1.0\" encoding=\"UTF-8\"?> 
<SOAP-ENV:Envelope
    xmlns:SOAP-ENV=\"http://schemas.xmlsoap.org/soap/envelope/\"
    xmlns:ns1=\"https://plazavip.clarodrive.com/dla/soap/\"
    xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\"
    xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"
    xmlns:SOAP-ENC=\"http://schemas.xmlsoap.org/soap/encoding/\" SOAP-ENV:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">
    <SOAP-ENV:Body>
        <ns1:actualizaCliente>
            <info xsi:type=\"xsd:string\">
                <peticion>
                    <id>
                        <![CDATA[%s]]>
                    </id>
                    <nombre>
                        <![CDATA[%s]]>
                    </nombre>
                    <apellidoPaterno>
                        <![CDATA[%s]]>
                    </apellidoPaterno>
                    <email>
                        <![CDATA[%s]]>
                    </email>
                    <telefonoTelmex>
                        <![CDATA[3133263499]]>
                    </telefonoTelmex>
                    <formaPago>
                        <![CDATA[2]]>
                    </formaPago>
                </peticion>
            </info>
        </ns1:actualizaCliente>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

// AdicionarClienteBody used on the add client
const AdicionarClienteBody = `<?xml version=\"1.0\" encoding=\"UTF-8\"?>  
<SOAP-ENV:Envelope
    xmlns:SOAP-ENV=\"http://schemas.xmlsoap.org/soap/envelope/\"
    xmlns:ns1=\"https://plazavip.clarodrive.com/dla/soap/\"
    xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\"
    xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"
    xmlns:SOAP-ENC=\"http://schemas.xmlsoap.org/soap/encoding/\" SOAP-ENV:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">
    <SOAP-ENV:Body>
        <ns1:altaCliente>
            <info xsi:type=\"xsd:string\">
                <peticion>
                    <id>
                        <![CDATA[%s]]>
                    </id> 
                    <nombre>
                        <![CDATA[%s]]>
                    </nombre> 
                    <apellidoPaterno>
                        <![CDATA[%s]]>
                    </apellidoPaterno> 
                    <email>
                        <![CDATA[%s]]>
                    </email> 
                    <telefonoTelmex>
                        <![CDATA[%s]]>
                    </telefonoTelmex> 
                    <formaPago>
                        <![CDATA[%s]]>
                    </formaPago> 
                </peticion>
            </info>
        </ns1:altaCliente>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

// BorrarClienteBody used on the borrar cliente
const BorrarClienteBody = `<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:soap="http://dla.d.iliux.com/dla/soap/">
<soapenv:Header/>
<soapenv:Body>
   <soap:borrarCliente soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
	  <id xsi:type="xsd:int">%s</id>
   </soap:borrarCliente>
</soapenv:Body>
</soapenv:Envelope>`

// AddSubscriptionBody used on the add subscription
const AddSubscriptionBody = `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" xmlns:ns1="https://plazavip.clarodrive.com/dla/soap/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
   <SOAP-ENV:Body>
      <ns1:altaSuscripcionCliente>
         <info xsi:type="xsd:string">
            <peticion>
               <idSuscripcion><![CDATA[%s]]></idSuscripcion>
               <nombre><![CDATA[%s]]></nombre>
               <item><![CDATA[%s]]></item>
               <leyenda><![CDATA[%s]]></leyenda>
               <idCliente><![CDATA[%s]]></idCliente>
               <fechaInicio><![CDATA[%s]]></fechaInicio>
               <fechaInicioCiclo><![CDATA[%s]]></fechaInicioCiclo>
               <idPago><![CDATA[%s]]></idPago>
               <billingPeriodo><![CDATA[MES]]></billingPeriodo>
               <billingFrecuencia><![CDATA[1]]></billingFrecuencia>
               <billingCiclo><![CDATA[0]]></billingCiclo>
               <billingPrecio><![CDATA[0]]></billingPrecio>
               <codigoPromo><![CDATA[%s]]></codigoPromo>
               <trialPeriodo><![CDATA[MES]]></trialPeriodo>
               <trialFrecuencia><![CDATA[0]]></trialFrecuencia>
               <trialCiclo><![CDATA[0]]></trialCiclo>
               <trialPrecio><![CDATA[0]]></trialPrecio>
            </peticion>
         </info>
      </ns1:altaSuscripcionCliente>
   </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

// CancelarSubscription is the body for removing subscription
const CancelarSubscription = `<?xml version=\"1.0\" encoding=\"UTF-8\"?> 
<SOAP-ENV:Envelope
    xmlns:SOAP-ENV=\"http://schemas.xmlsoap.org/soap/envelope/\"
    xmlns:ns1=\"https://plazavip.clarodrive.com/dla/soap/\"
    xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\"
    xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"
    xmlns:SOAP-ENC=\"http://schemas.xmlsoap.org/soap/encoding/\" SOAP-ENV:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">
    <SOAP-ENV:Body>
        <ns1:cancelaSuscripcionCliente>
            <info xsi:type=\"xsd:string\">
                <peticion>
                    <idCliente>
                        <![CDATA[%s]]>
                    </idCliente>
                    <idSuscripcion>
                        <![CDATA[%s]]>
                    </idSuscripcion>
                    <fechaFin>
                        <![CDATA[%s]]>
                    </fechaFin>
                </peticion>
            </info>
        </ns1:cancelaSuscripcionCliente>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`
