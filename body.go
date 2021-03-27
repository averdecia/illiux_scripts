package main

// ObtenerClienteBody used on the show client request
const ObtenerClienteBody = `<?xml version="1.0" encoding="UTF-8"?> 
<SOAP-ENV:Envelope
    xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/"
    xmlns:ns1="https://plazavip.clarodrive.com/dla/soap/"
    xmlns:xsd="http://www.w3.org/2001/XMLSchema"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
    <SOAP-ENV:Body>
        <ns1:obtenerCliente>
            <id xsi:type="xsd:int">%s</id>
        </ns1:obtenerCliente>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

// ActualizarClienteBody used on the update client request
const ActualizarClienteBody = `<?xml version="1.0" encoding="UTF-8"?> 
<SOAP-ENV:Envelope
    xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/"
    xmlns:ns1="https://plazavip.clarodrive.com/dla/soap/"
    xmlns:xsd="http://www.w3.org/2001/XMLSchema"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
    <SOAP-ENV:Body>
        <ns1:actualizaCliente>
        <info xsi:type="xsd:string">
        &lt;peticion&gt;
        &lt;id&gt;%s&lt;/id&gt;
        &lt;nombre&gt;%s&lt;/nombre&gt;
        &lt;apellidoPaterno&gt;%s&lt;/apellidoPaterno&gt;
        &lt;email&gt;%s&lt;/email&gt;
        &lt;telefonoTelmex&gt;%s&lt;/telefonoTelmex&gt;
        &lt;formaPago&gt;%s&lt;/formaPago&gt;
        &lt;/peticion&gt;
        </info>
        </ns1:actualizaCliente>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

// AdicionarClienteBody used on the add client
const AdicionarClienteBody = `<?xml version="1.0" encoding="UTF-8"?>  
<SOAP-ENV:Envelope
    xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/"
    xmlns:ns1="https://plazavip.clarodrive.com/dla/soap/"
    xmlns:xsd="http://www.w3.org/2001/XMLSchema"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
    <SOAP-ENV:Body>
        <ns1:altaCliente>
        <info xsi:type="xsd:string">
        &lt;peticion&gt;
        &lt;id&gt;%s&lt;/id&gt;
        &lt;nombre&gt;%s&lt;/nombre&gt;
        &lt;apellidoPaterno&gt;%s&lt;/apellidoPaterno&gt;
        &lt;email&gt;%s&lt;/email&gt;
        &lt;telefonoTelmex&gt;%s&lt;/telefonoTelmex&gt;
        &lt;formaPago&gt;%s&lt;/formaPago&gt;
        &lt;/peticion&gt;
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
         &lt;peticion&gt;
         &lt;idSuscripcion&gt;&lt;![CDATA[%s]]&gt;&lt;/idSuscripcion&gt;
         &lt;nombre&gt;&lt;![CDATA[%s]]&gt;&lt;/nombre&gt;
         &lt;item&gt;&lt;![CDATA[%s]]&gt;&lt;/item&gt;
         &lt;leyenda&gt;&lt;![CDATA[%s]]&gt;&lt;/leyenda&gt;
         &lt;idCliente&gt;&lt;![CDATA[%s]]&gt;&lt;/idCliente&gt;
         &lt;fechaInicio&gt;&lt;![CDATA[%s]]&gt;&lt;/fechaInicio&gt;
         &lt;fechaInicioCiclo&gt;&lt;![CDATA[%s]]&gt;&lt;/fechaInicioCiclo&gt;
         &lt;idPago&gt;&lt;![CDATA[%s]]&gt;&lt;/idPago&gt;
         &lt;billingPeriodo&gt;&lt;![CDATA[MES]]&gt;&lt;/billingPeriodo&gt;
         &lt;billingFrecuencia&gt;&lt;![CDATA[1]]&gt;&lt;/billingFrecuencia&gt;
         &lt;billingCiclo&gt;&lt;![CDATA[0]]&gt;&lt;/billingCiclo&gt;
         &lt;billingPrecio&gt;&lt;![CDATA[0]]&gt;&lt;/billingPrecio&gt;
         &lt;codigoPromo&gt;&lt;![CDATA[%s]]&gt;&lt;/codigoPromo&gt;
         &lt;trialPeriodo&gt;&lt;![CDATA[MES]]&gt;&lt;/trialPeriodo&gt;
         &lt;trialFrecuencia&gt;&lt;![CDATA[0]]&gt;&lt;/trialFrecuencia&gt;
         &lt;trialCiclo&gt;&lt;![CDATA[0]]&gt;&lt;/trialCiclo&gt;
         &lt;trialPrecio&gt;&lt;![CDATA[0]]&gt;&lt;/trialPrecio&gt;
         &lt;/peticion&gt;
         </info>
      </ns1:altaSuscripcionCliente>
   </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

// CancelarSubscription is the body for removing subscription
const CancelarSubscription = `<?xml version="1.0" encoding="UTF-8"?> 
<SOAP-ENV:Envelope
    xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/"
    xmlns:ns1="https://plazavip.clarodrive.com/dla/soap/"
    xmlns:xsd="http://www.w3.org/2001/XMLSchema"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
    <SOAP-ENV:Body>
        <ns1:cancelaSuscripcionCliente>
        <info xsi:type="xsd:string">
        &lt;peticion&gt;
                   &lt;idCliente&gt;
                       &lt;![CDATA[%s]]&gt;
                   &lt;/idCliente&gt;
                   &lt;idSuscripcion&gt;
                       &lt;![CDATA[%s]]&gt;
                   &lt;/idSuscripcion&gt;
               &lt;/peticion&gt;
        </info>
        </ns1:cancelaSuscripcionCliente>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

const SendEmailBody = `{
	"service": "CLARODRIVE",
	"region": "MX",
	"type": "%s",
	"email": "%s",
	"user_id": "%s",
	"variables": {
			"email": "%s",
			"last_active_date": "26/05/2021"
		},
	"extras": {
		"subject": "Actualiza tu información de pago"
	},
	"dryrun": false
}`
