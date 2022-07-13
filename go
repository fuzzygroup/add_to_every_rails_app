#!/bin/bash

ECHO "execute this from the root directory of a freshly created Rails app"
read -p "Press any key to continue or CTRL+C to abort"
mkdir docs
mkdir app/views/pages
cp -p ../add_to_every_rails_app/docs/generic_class.txt docs
cp -p ../add_to_every_rails_app/pages/* app/views/pages
cp -p ../add_to_every_rails_app/controllers/pages_controller.rb app/views/controllers
cp -p ../add_to_every_rails_app/bin/ssh* bin
cp -p ../add_to_every_rails_app/db/migrate/* db/migrate
cp -p ../add_to_every_rails_app/pre-commit .git/hooks
bundle add url_common
bundle add text_common
bundle add date_common
bundle add time_common 
bundle add awesome_print
bundle add factory_bot_rails --group test
bundle add solargraph --group development

ECHO "ADD THIS TO ROUTES:"
ECHO "========================================================================"
ECHO "match '/pages/', via: :all,           to: 'pages#index'"
ECHO "match '/pages/about', via: :all,      to: 'pages#about'"
ECHO "match '/pages/contact', via: :all,    to: 'pages#contact'"
ECHO "match '/pages/dedication', via: :all, to: 'pages#dedication'"
ECHO "match '/pages/faq', via: :all,        to: 'pages#faq'"
ECHO "match '/pages/jobs', via: :all,       to: 'pages#jobs'"
ECHO "match '/pages/mission', via: :all,    to: 'pages#mission'"
ECHO "match '/pages/values', via: :all,     to: 'pages#values'"

ECHO "NOW FIX THE PORT VALUE incrementing up from 5000 by 100s"