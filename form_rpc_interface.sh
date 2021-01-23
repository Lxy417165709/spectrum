cat template_of_rpc_interface.go > tmp_template_of_rpc_interface.go
rpl 'DeleteFavorForGood' "$1" tmp_template_of_rpc_interface.go
cat tmp_template_of_rpc_interface.go
