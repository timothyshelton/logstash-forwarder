#!/usr/bin/env ruby
#

require "json"

class LogFile
  ISO8601_STRFTIME = "%04d-%02d-%02dT%02d:%02d:%02d.%06d%+03d:00"

  def initialize(path)
    @path = path
    @count = 0
    open
  end

  def open
    @file.close if !@file.nil?
    @file = File.new(@path, "w+")
  end

  def rotate
    File.rename(@path, "#{@path}.old")
    open
  end

  def hit
    @file.syswrite({
      "count" => @count,
      "time" => iso8601(Time.now)
    }.to_json)
    @count += 1
  end

  def iso8601(t)
    return sprintf(ISO8601_STRFTIME, t.year, t.month, t.day, t.hour, t.min,
                   t.sec, t.tv_usec, t.utc_offset / 3600)
  end
end

path = ARGV[0]

numlogs = 50
logs = numlogs.times.collect { |i| LogFile.new(File.join(path, "#{i}.log")) }

rotate_count = 50000

count = 0
while true
  hits = (numlogs * rand).to_i
  hits.times do |i|
    logs[i].hit
  end

  count += 1
  if count > rotate_count
    # rotate 5 logs at random.
    5.times do 
      logs[(numlogs * rand).to_i].rotate
    end
    count = 0
  end

  sleep(rand * 0.1)
end