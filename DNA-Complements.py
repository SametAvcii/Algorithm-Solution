def DNA_strand(dna):
 
    output =""
    for i in dna:
        if i=="A":
            output+="T"
        elif i=="T":    
            output+="A"
        elif i=="G":
            output+="C"
        elif i=="C":
            output+="G"
          
          
    return output    
  