#!/usr/bin/env ruby

require 'securerandom'
require 'json'

recinfos = []
(102..255).each do |i|
  recinfos << [SecureRandom.uuid, "dns", "z1", "default", "cf", "127.0.0.#{i}", "bosh", SecureRandom.uuid]
end

(1..8).each do |i|
  (1..255).each do |j|
    recinfos << [SecureRandom.uuid, "dns", "z1", "default", "cf", "127.0.#{i}.#{j}", "bosh", SecureRandom.uuid]
  end
end

json_hash = {}

json_hash['records'] = recinfos.map{|info| [info[5], "#{info[0]}.#{info[1]}.#{info[3]}.#{info[4]}.#{info[6]}"]}
json_hash['version'] = 11235
json_hash['record_keys'] = ['id','instance_group','az','network','deployment','ip','domain','agent_id']
json_hash['record_infos'] = recinfos

puts JSON.pretty_generate(json_hash)
