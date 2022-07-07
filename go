#!/bin/bash

ECHO "execute this from the root directory of a freshly created Rails app"
read -p "Press any key to continue or CTRL+C to abort"
mkdir docs
cp -p ../add_scott_to_new_rails_app/docs/generic_class.txt docs
cp -p ../add_scott_to_new_rails_app/bin/ssh* bin
cp -p ../add_scott_to_new_rails_app/db/migrate/* db/migrate
bundle add url_common
bundle add text_common
bundle add date_common
bundle add time_common 
bundle add awesome_print
bundle add factory_bot_rails --group test

