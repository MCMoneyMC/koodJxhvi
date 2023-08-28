id=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d '#' -f2 )
echo $id
cat interviews/interview-"$id"
echo $MAIN_SUSPECT