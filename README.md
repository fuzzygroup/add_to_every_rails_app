# add_to_every_rails_app

This is a shell script which I run AFTER a new rails app is created and which makes said app more like how I like it.

I build a lot of test rails apps.  Like a lot.  Here's an example:

    ❯ mdfind -name application_controller.rb | wc -l
         714

Sure some of those are dupes when DropBox when insane to the brain earlier this year but even if the number is 100 then its still a lot.  And because I want my applications to generally function the same, I find myself doing the same configuration crap over and over and over and over.  And now I have a shell script which does that for me.

**NOTE**: WHAT YOU WANT TO DO IS NOT USE THIS DIRECTLY; YOU AREN'T ME.  Fork it to your repo and adjust it.

Assuming you do your work in a /rails/ directory which holds multiple apps:

1. Clone this locally into /rails/: git clone git@github.com:fuzzygroup/add_to_every_rails_app.git
2. In your application directory -- say foo -- so you are in /rails/foo, run the **go** shell command.

Here's what this gives for output:

    ❯ ../add_to_every_rails_app/go
    execute this from the root directory of a freshly created Rails app
    Press any key to continue or CTRL+C to abort
    mkdir: docs: File exists
    cp: app/views/pages is not a directory
    Your Gemfile lists the gem factory_bot_rails (~> 6.2) more than once.
    You should probably keep only one of them.
    Remove any duplicate entries and specify the gem only once.
    While it's not a problem now, it could cause errors if you change the version of one of them later.

    [!] There was an error parsing `injected gems`: You cannot specify the same gem twice with different version requirements.
    You specified: url_common (~> 0.1.1) and url_common (>= 0). Gem already added. Bundler cannot continue.

     #  from injected gems:1
     #  -------------------------------------------
     >  gem "url_common", ">= 0"
     #  -------------------------------------------
    Your Gemfile lists the gem factory_bot_rails (~> 6.2) more than once.
    You should probably keep only one of them.
    Remove any duplicate entries and specify the gem only once.
    While it's not a problem now, it could cause errors if you change the version of one of them later.

    [!] There was an error parsing `injected gems`: You cannot specify the same gem twice with different version requirements.
    You specified: text_common (~> 0.1.0) and text_common (>= 0). Gem already added. Bundler cannot continue.

     #  from injected gems:1
     #  -------------------------------------------
     >  gem "text_common", ">= 0"
     #  -------------------------------------------
    Your Gemfile lists the gem factory_bot_rails (~> 6.2) more than once.
    You should probably keep only one of them.
    Remove any duplicate entries and specify the gem only once.
    While it's not a problem now, it could cause errors if you change the version of one of them later.

    [!] There was an error parsing `injected gems`: You cannot specify the same gem twice with different version requirements.
    You specified: date_common (~> 0.13) and date_common (>= 0). Gem already added. Bundler cannot continue.

     #  from injected gems:1
     #  -------------------------------------------
     >  gem "date_common", ">= 0"
     #  -------------------------------------------
    Your Gemfile lists the gem factory_bot_rails (~> 6.2) more than once.
    You should probably keep only one of them.
    Remove any duplicate entries and specify the gem only once.
    While it's not a problem now, it could cause errors if you change the version of one of them later.

    [!] There was an error parsing `injected gems`: You cannot specify the same gem twice with different version requirements.
    You specified: time_common (~> 0.1.1) and time_common (>= 0). Gem already added. Bundler cannot continue.

     #  from injected gems:1
     #  -------------------------------------------
     >  gem "time_common", ">= 0"
     #  -------------------------------------------
    Your Gemfile lists the gem factory_bot_rails (~> 6.2) more than once.
    You should probably keep only one of them.
    Remove any duplicate entries and specify the gem only once.
    While it's not a problem now, it could cause errors if you change the version of one of them later.
    Fetching gem metadata from https://rubygems.org/.........
    Resolving dependencies..............
    Fetching gem metadata from https://rubygems.org/.........
    Resolving dependencies...........
    Using rake 13.0.6
    Using builder 3.2.4
    Using erubi 1.10.0
    Using crass 1.0.6
    Using digest 3.1.0
    Using bundler 2.3.16
    Using concurrent-ruby 1.1.10
    Using kaminari-core 1.2.2
    Using ffi 1.15.5
    Using minitest 5.16.0
    Using selectize-rails 0.12.6
    Using nio4r 2.5.8
    Using marcel 1.0.2
    Using smart_properties 1.17.0
    Using timeout 0.3.0
    Using strscan 3.0.3
    Using matrix 0.4.2
    Using regexp_parser 2.5.0
    Using childprocess 4.1.0
    Using chronic 0.10.2
    Using chunky_png 1.4.0
    Using will_paginate 3.3.1
    Using connection_pool 2.2.5
    Using msgpack 1.5.2
    Using io-console 0.5.11
    Using orm_adapter 0.5.0
    Using unf_ext 0.0.8.2
    Using rainbow 3.1.1
    Using parallel 1.22.1
    Using rexml 3.2.5
    Using ruby-progressbar 1.11.0
    Using unicode-display_width 2.1.0
    Using thor 1.2.1
    Using racc 1.6.0
    Using public_suffix 4.0.7
    Using hashie 5.0.0
    Using rack 2.2.3.1
    Using tilt 2.0.10
    Using ast 2.4.2
    Using websocket-extensions 0.1.5
    Using date_common 0.13
    Using mini_magick 4.11.0
    Using multi_xml 0.6.0
    Using iniparse 1.5.0
    Using json 2.6.2
    Using local_time 2.1.0
    Using net-http-digest_auth 1.4.1
    Using rubyntlm 0.6.3
    Using webrick 1.7.0
    Using oj 3.13.14
    Using pg 1.4.0
    Using redcarpet 3.5.1
    Using honeybadger 4.12.1
    Using redis 4.6.0
    Using awesome_print 1.9.2
    Using pdf-core 0.9.0
    Using ruby-oembed 0.16.1
    Using mini_mime 1.1.2
    Using websocket 1.2.9
    Using text_common 0.1.0
    Using time_common 0.1.1
    Using i18n 1.10.0
    Using stripe 6.4.0
    Using webrobots 0.1.2
    Using ttfunk 1.7.0
    Using http-form_data 2.3.0
    Using mime-types-data 3.2022.0105
    Using html_tokenizer 0.0.7
    Using rotp 6.2.0
    Using whenever 1.0.0
    Using rubyzip 2.3.2
    Using bcrypt 3.1.18
    Using bootsnap 1.12.0
    Using reline 0.3.1
    Using fuzzyurl 0.9.0
    Using hashids 1.0.6
    Using zeitwerk 2.6.0
    Using tzinfo 2.0.4
    Using sprockets 4.0.3
    Using rack-protection 2.2.0
    Using parser 3.1.2.0
    Using websocket-driver 0.7.5
    Using net-protocol 0.1.3
    Using net-http-persistent 4.0.1
    Using rqrcode_core 1.2.0
    Using unf 0.1.4
    Using nokogiri 1.13.6 (x86_64-darwin)
    Using bindex 0.8.1
    Using method_source 1.0.0
    Using irb 1.4.1
    Using rubocop-ast 1.18.0
    Using request_store 1.5.1
    Using sassc 2.4.0
    Using omniauth 2.1.0
    Using tzinfo-data 1.2022.1
    Using net-imap 0.2.3
    Using net-smtp 0.3.1
    Using overcommit 0.59.1
    Using sidekiq 6.5.1
    Using mail 2.7.1
    Using mime-types 3.4.1
    Using debug 1.5.0
    Using addressable 2.8.0
    Using httparty 0.20.0
    Using rack-test 1.1.0
    Using launchy 2.5.0
    Using warden 1.2.9
    Using meilisearch 0.18.3
    Using ruby-vips 2.1.4
    Using meilisearch-rails 0.5.1 from https://github.com/meilisearch/meilisearch-rails.git (at release-v0.5.1@f1eb9cc)
    Using puma 5.6.4
    Using image_processing 1.12.2
    Using domain_name 0.5.20190701
    Using prawn 2.4.0 from https://github.com/prawnpdf/prawn.git (at master@7d4f6b8)
    Using loofah 2.18.0
    Using http-cookie 1.0.5
    Using rubocop 1.29.1
    Using xpath 3.2.0
    Using selenium-webdriver 4.2.1
    Using capybara 3.37.1
    Using mechanize 2.8.5
    Using rubocop-performance 1.13.3
    Using webdrivers 5.0.0
    Using letter_opener 1.8.1
    Using rails-html-sanitizer 1.4.3
    Using ffi-compiler 1.0.1
    Using rqrcode 2.1.1
    Using llhttp-ffi 0.4.0
    Using net-pop 0.1.1
    Using standard 1.12.1
    Using url_common 0.1.2
    Using activesupport 7.0.3
    Using http 5.1.0
    Using rails-dom-testing 2.0.3
    Using globalid 1.0.0
    Using activemodel 7.0.3
    Using factory_bot 6.2.1
    Using inline_svg 1.8.0
    Using actionview 7.0.3
    Using pagy 5.10.1
    Using name_of_person 1.1.1
    Using actionpack 7.0.3
    Using activerecord 7.0.3
    Using kaminari-actionview 1.2.2
    Using better_html 1.0.16
    Using jbuilder 2.11.5 from https://github.com/excid3/jbuilder.git (at partial-paths@394014a)
    Using pundit 2.2.0
    Using activejob 7.0.3
    Using actioncable 7.0.3
    Using acts-as-taggable-on 9.0.1
    Using railties 7.0.3
    Using kaminari-activerecord 1.2.2
    Using annotate 3.1.1 from https://github.com/excid3/annotate_models.git (at rails7@34ab955)
    Using erb_lint 0.1.3
    Using omniauth-rails_csrf_protection 1.0.1
    Using pg_search 2.3.6
    Using pretender 0.4.0
    Using sprockets-rails 3.4.2
    Using activestorage 7.0.3
    Using jquery-rails 4.5.0
    Using kaminari 1.2.2
    Using cssbundling-rails 1.1.1
    Using responders 3.0.1 from https://github.com/excid3/responders.git (at fix-redirect-status-config@1288c53)
    Using factory_bot_rails 6.2.0
    Using actionmailbox 7.0.3
    Using actiontext 7.0.3
    Using jsbundling-rails 1.0.3
    Using stimulus-rails 1.0.4
    Using turbo-rails 1.1.1
    Using sassc-rails 2.1.2
    Using web-console 4.2.0
    Using actionmailer 7.0.3
    Using administrate 0.17.0 from https://github.com/excid3/administrate.git (at jumpstart@238d6d5)
    Using rails 7.0.3
    Using letter_opener_web 2.0.0
    Using jumpstart 0.1.0 from source at `lib/jumpstart`
    Using noticed 1.5.9
    Using acts_as_tenant 0.5.1
    Using commontator 7.0.0
    Using invisible_captcha 2.0.0
    Using pay 3.0.24
    Using prefixed_ids 1.2.2
    Using prawn-table 0.2.2
    Using receipts 1.1.2
    Using devise 4.8.1
    Using devise-i18n 1.10.2
    Using administrate-field-active_storage 0.4.1
    Your Gemfile lists the gem factory_bot_rails (~> 6.2) more than once.
    You should probably keep only one of them.
    Remove any duplicate entries and specify the gem only once.
    While it's not a problem now, it could cause errors if you change the version of one of them later.

    [!] There was an error parsing `injected gems`: You cannot specify the same gem twice with different version requirements.
    You specified: factory_bot_rails (~> 6.2) and factory_bot_rails (>= 0). Gem already added. Bundler cannot continue.

     #  from injected gems:1
     #  -------------------------------------------
     >  gem "factory_bot_rails", ">= 0", :group => :test
     #  -------------------------------------------
    Your Gemfile lists the gem factory_bot_rails (~> 6.2) more than once.
    You should probably keep only one of them.
    Remove any duplicate entries and specify the gem only once.
    While it's not a problem now, it could cause errors if you change the version of one of them later.
    Fetching gem metadata from https://rubygems.org/.........
    Resolving dependencies.............
    Fetching gem metadata from https://rubygems.org/.........
    Resolving dependencies............................
    Using rake 13.0.6
    Using concurrent-ruby 1.1.10
    Using erubi 1.10.0
    Using zeitwerk 2.6.0
    Using method_source 1.0.0
    Using kaminari-core 1.2.2
    Using rack 2.2.3.1
    Using public_suffix 4.0.7
    Using ffi 1.15.5
    Using tilt 2.0.10
    Using selectize-rails 0.12.6
    Using ast 2.4.2
    Using awesome_print 1.9.2
    Using backport 1.2.0
    Using bcrypt 3.1.18
    Using benchmark 0.2.0
    Using html_tokenizer 0.0.7
    Using smart_properties 1.17.0
    Using bindex 0.8.1
    Using regexp_parser 2.5.0
    Using nio4r 2.5.8
    Using crass 1.0.6
    Using childprocess 4.1.0
    Using chronic 0.10.2
    Using chunky_png 1.4.0
    Using will_paginate 3.3.1
    Using msgpack 1.5.2
    Using date_common 0.13
    Using io-console 0.5.11
    Using orm_adapter 0.5.0
    Using diff-lcs 1.5.0
    Using unf_ext 0.0.8.2
    Using builder 3.2.4
    Using racc 1.6.0
    Using rexml 3.2.5
    Using e2mmap 0.1.0
    Using rainbow 3.1.1
    Using websocket-extensions 0.1.5
    Using marcel 1.0.2
    Using honeybadger 4.12.1
    Using bundler 2.3.16
    Using connection_pool 2.2.5
    Using mini_mime 1.1.2
    Using timeout 0.3.0
    Using iniparse 1.5.0
    Using jaro_winkler 1.5.4
    Using json 2.6.2
    Using local_time 2.1.0
    Using net-http-digest_auth 1.4.1
    Using mini_magick 4.11.0
    Using webrick 1.7.0
    Using webrobots 0.1.2
    Using hashids 1.0.6
    Using fuzzyurl 0.9.0
    Using pdf-core 0.9.0
    Using pg 1.4.0
    Using ttfunk 1.7.0
    Using redcarpet 3.5.1
    Using redis 4.6.0
    Using oj 3.13.14
    Using rqrcode_core 1.2.0
    Using ruby-oembed 0.16.1
    Using websocket 1.2.9
    Using stripe 6.4.0
    Using text_common 0.1.0
    Using hashie 5.0.0
    Using digest 3.1.0
    Using http-form_data 2.3.0
    Using addressable 2.8.0
    Using sassc 2.4.0
    Using parser 3.1.2.0
    Using ffi-compiler 1.0.1
    Using ruby-vips 2.1.4
    Using rack-test 1.1.0
    Using request_store 1.5.1
    Using sprockets 4.0.3
    Using warden 1.2.9
    Using rack-protection 2.2.0
    Using bootsnap 1.12.0
    Using reline 0.3.1
    Using tzinfo 2.0.4
    Using whenever 1.0.0
    Using i18n 1.10.0
    Using nokogiri 1.13.6 (x86_64-darwin)
    Using multi_xml 0.6.0
    Using mime-types-data 3.2022.0105
    Using kramdown 2.4.0
    Using websocket-driver 0.7.5
    Using net-http-persistent 4.0.1
    Using mail 2.7.1
    Using net-protocol 0.1.3
    Using overcommit 0.59.1
    Using sidekiq 6.5.1
    Using yard 0.9.28
    Using rubocop-ast 1.18.0
    Using unicode-display_width 2.1.0
    Using llhttp-ffi 0.4.0
    Using image_processing 1.12.2
    Using omniauth 2.1.0
    Using rubyntlm 0.6.3
    Using tzinfo-data 1.2022.1
    Using net-pop 0.1.1
    Using net-smtp 0.3.1
    Using mime-types 3.4.1
    Using kramdown-parser-gfm 1.1.0
    Using loofah 2.18.0
    Using rubyzip 2.3.2
    Using matrix 0.4.2
    Using rails-html-sanitizer 1.4.3
    Using puma 5.6.4
    Using launchy 2.5.0
    Using time_common 0.1.1
    Using letter_opener 1.8.1
    Using rqrcode 2.1.1
    Using strscan 3.0.3
    Using irb 1.4.1
    Using rotp 6.2.0
    Using net-imap 0.2.3
    Using minitest 5.16.0
    Using parallel 1.22.1
    Using xpath 3.2.0
    Using reverse_markdown 2.1.1
    Using httparty 0.20.0
    Using ruby-progressbar 1.11.0
    Using selenium-webdriver 4.2.1
    Using rubocop 1.29.1
    Using webdrivers 5.0.0
    Using debug 1.5.0
    Using rubocop-performance 1.13.3
    Using activesupport 7.0.3
    Using rails-dom-testing 2.0.3
    Using factory_bot 6.2.1
    Using thor 1.2.1
    Using globalid 1.0.0
    Using capybara 3.37.1
    Using activejob 7.0.3
    Using solargraph 0.45.0
    Using name_of_person 1.1.1
    Using pagy 5.10.1
    Using pundit 2.2.0
    Using standard 1.12.1
    Using meilisearch 0.18.3
    Using prawn 2.4.0 from https://github.com/prawnpdf/prawn.git (at master@7d4f6b8)
    Using activemodel 7.0.3
    Using meilisearch-rails 0.5.1 from https://github.com/meilisearch/meilisearch-rails.git (at release-v0.5.1@f1eb9cc)
    Using actionview 7.0.3
    Using inline_svg 1.8.0
    Using unf 0.1.4
    Using activerecord 7.0.3
    Using actionpack 7.0.3
    Using kaminari-actionview 1.2.2
    Using actioncable 7.0.3
    Using pretender 0.4.0
    Using jbuilder 2.11.5 from https://github.com/excid3/jbuilder.git (at partial-paths@394014a)
    Using activestorage 7.0.3
    Using domain_name 0.5.20190701
    Using actiontext 7.0.3
    Using railties 7.0.3
    Using acts-as-taggable-on 9.0.1
    Using cssbundling-rails 1.1.1
    Using jsbundling-rails 1.0.3
    Using stimulus-rails 1.0.4
    Using turbo-rails 1.1.1
    Using web-console 4.2.0
    Using pg_search 2.3.6
    Using better_html 1.0.16
    Using actionmailer 7.0.3
    Using actionmailbox 7.0.3
    Using erb_lint 0.1.3
    Using jquery-rails 4.5.0
    Using kaminari-activerecord 1.2.2
    Using sprockets-rails 3.4.2
    Using letter_opener_web 2.0.0
    Using annotate 3.1.1 from https://github.com/excid3/annotate_models.git (at rails7@34ab955)
    Using kaminari 1.2.2
    Using sassc-rails 2.1.2
    Using omniauth-rails_csrf_protection 1.0.1
    Using rails 7.0.3
    Using responders 3.0.1 from https://github.com/excid3/responders.git (at fix-redirect-status-config@1288c53)
    Using factory_bot_rails 6.2.0
    Using acts_as_tenant 0.5.1
    Using administrate 0.17.0 from https://github.com/excid3/administrate.git (at jumpstart@238d6d5)
    Using invisible_captcha 2.0.0
    Using jumpstart 0.1.0 from source at `lib/jumpstart`
    Using commontator 7.0.0
    Using pay 3.0.24
    Using prefixed_ids 1.2.2
    Using http-cookie 1.0.5
    Using http 5.1.0
    Using mechanize 2.8.5
    Using noticed 1.5.9
    Using url_common 0.1.2
    Using prawn-table 0.2.2
    Using receipts 1.1.2
    Using administrate-field-active_storage 0.4.1
    Using devise 4.8.1
    Using devise-i18n 1.10.2
    ADD THIS TO ROUTES:
    ========================================================================
    match '/pages/', via: :all,           to: 'pages#index'
    match '/pages/about', via: :all,      to: 'pages#about'
    match '/pages/contact', via: :all,    to: 'pages#contact'
    match '/pages/dedication', via: :all, to: 'pages#dedication'
    match '/pages/faq', via: :all,        to: 'pages#faq'
    match '/pages/jobs', via: :all,       to: 'pages#jobs'
    match '/pages/mission', via: :all,    to: 'pages#mission'
    match '/pages/values', via: :all,     to: 'pages#values'
    NOW FIX THE PORT VALUE incrementing up from 5000 by 100s
    


