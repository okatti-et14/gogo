package main

import "fmt"

func main() {
	SelectCustListQueryWhereFreeWord := `
	AND (
		lastname || firstname ILIKE %v 
		OR lastnamekana || firstnamekana ILIKE %v 
		OR zipcd LIKE %v  
		OR (pc.kbnnm || address1 || address2 || address3 ILIKE  %[1]v)  
		OR EXISTS (SELECT * FROM mcusttelno ct WHERE c.custid = ct.custid AND ct.telno LIKE %[1]v) 
		OR EXISTS (SELECT * FROM mcustmailaddress cm WHERE c.custid = cm.custid AND cm.mailaddress ILIKE %[0]v) 
		OR workzipcd LIKE %v  OR (pw.kbnnm || workaddress1 || workaddress2 || workaddress3) ILIKE %v  
		OR EXISTS (SELECT * FROM mcustworktelno wt WHERE c.custid = wt.custid AND wt.telno LIKE %v) 
		OR EXISTS (SELECT * FROM mcustworkmailaddress wm WHERE c.custid = wm.custid AND wm.mailaddress ILIKE %v) 
		OR workplacename ILIKE %v 
		OR workplacenamekana ILIKE %v 
		OR workdivision ILIKE %v  
		OR workpost ILIKE %v 
		OR c.custnote1 ILIKE %v 
		OR custnote2 ILIKE %v 
		OR custnote3 ILIKE %v 
		OR EXISTS (SELECT * FROM mcustshopcustnote sc WHERE c.custid = sc.custid AND sc.custnote1 ILIKE %v) 
		OR EXISTS (SELECT * FROM mcustmbrcode WHERE c.custid = mcustmbrcode.custid AND (mbrcodeno LIKE %v OR mbrcodememo ILIKE %v)) 
	)`
	fmt.Printf(SelectCustListQueryWhereFreeWord, "aa")
}
