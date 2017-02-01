# weerwolvensim
Simulator voor Weerwolven van Wakkerdam


#nacht
cupido assign
cupido execute       //mark two lovers

ziener assign
ziener execute       //do something

genezer assign
genezer execute		 //mark player as protected

weerwolven assign
weerwolven execute 	 //mark player as attacked

heks assign
heks execute		 //mark attacked player as revived, mark other player as poisoned

heks validate		 //revived? remove attacked, poisoned? set dead
genezer validate	 //protected? remove attacked

dorpsoudste assign
dorpsoudste execute  //attacked? remove attacked

weerwolven validate  //attacked? set dead
#nacht end

#dag
assign all dead
each dead :
   jager execute		//dead? mark player as shot
   jager validate	   	//shot? set dead
   dorpsoudste validate //shot or poisoned? all residents lose powers
   cupido validate		//dead? set dead

assign burgemeester

vote execute			//mark player as lynched
vote validate			//lynched? set dead

#dag end


