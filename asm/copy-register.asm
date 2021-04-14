load R1 ff
load R2 aa
copy R1 R2
copy @0001 @0002 // absolute <- absolute
copy @0001 #0002 // absolute <- offset
copy @0001 *0002 // absolute <- pointer
copy #0001 @0002 // offset <- absolute
copy #0001 #0002 // offset <- offset
copy #0001 *0002 // offset <- pointer
copy *0001 @0002 // pointer <- absolute
copy *0001 #0002 // pointer <- offset
copy *0001 *0002 // pointer <- pointer
STOP