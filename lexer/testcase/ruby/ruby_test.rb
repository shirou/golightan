require 'minitest/autorun'

class RubyTest < Minitest::Test
# https://bugs.ruby-lang.org/issues/10617 --------------
def some_method_returning_array_or_nil(flag)
  flag ? [1, 2] : nil
end

def test_multiple_assignment_in_conditional_expression
  if (a, b = some_method_returning_array_or_nil(true))
    assert true
  else
    flunk 'should be true'
  end
  if (a, b = some_method_returning_array_or_nil(false))
    flunk 'should be false'
  else
    assert true
  end
end
# https://bugs.ruby-lang.org/issues/10617 --------------
# Array#max and Array#min.
# https://bugs.ruby-lang.org/issues/12172
# No test.

# https://bugs.ruby-lang.org/issues/12217
def test_enumerable_sum
  assert_equal 10, [1, 2, 3, 4].sum
  assert_equal 15, [1, 2, 3, 4].sum(5)
  assert_equal 0, [].sum
  refute_equal 0.3, [0.1, 0.1, 0.1].sum
  assert_equal 'foobar', ['foo', 'bar'].sum('')
  assert_equal '>>foobar', ['foo', 'bar'].sum('>>')
  assert_equal [1, 2, 3, 1, 5], [[1, 2], [3, 1, 5]].sum([])
end

# https://bugs.ruby-lang.org/issues/10594
def test_comparable_clamp
  assert_equal 2, 2.clamp(1, 3)
  assert_equal 1, 0.clamp(1, 3)
  assert_equal -1, 0.clamp(-2, -1)
  assert_equal 'b', 'c'.clamp('a', 'b')
  assert_raises(ArgumentError) {
    0.clamp(2, 1)
  }
end

# https://bugs.ruby-lang.org/issues/10121
def test_dir_empty?
  Dir.mktmpdir do |dir|
    assert Dir.empty?(dir)
  end
end

# https://bugs.ruby-lang.org/issues/11090
def test_enumerable_uniq
  olimpics = {
      1896 => 'Athens',
      1900 => 'Paris',
      1904 => 'Chikago',
      1906 => 'Athens',
      1908 => 'Rome'
  }
  each_city_first_time = olimpics.uniq { |k, v| v }
  assert_equal [
      [1896, "Athens"],
      [1900, "Paris"],
      [1904, "Chikago"],
      [1908, "Rome"]
  ], each_city_first_time

  values = (1..Float::INFINITY).lazy.uniq { |x| (x**2) % 10 }.first(6)
  assert_equal [1, 2, 3, 4, 5, 10], values
  # [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
  # [1, 4, 9, 6, 5, 6, 9, 4, 1, 0]
end

# https://bugs.ruby-lang.org/issues/9969
def test_file_empty?
  require 'tempfile'
  Tempfile.create("foo") do |f|
    assert File.empty?(f)
    assert File.zero?(f)
  end
end

# https://bugs.ruby-lang.org/issues/12245
def test_ceil_floor_truncate
  # 切り上げ
  assert_equal 2, 1.11.ceil
  assert_equal -1, -1.11.ceil
  assert_equal 1.2, 1.11.ceil(1)
  assert_equal -1.1, -1.11.ceil(1)
  # 負の値も指定可能
  assert_equal 11120, 11111.ceil(-1)
  assert_equal -11110, -11111.ceil(-1)

  # 切り下げ
  assert_equal 1, 1.99.floor
  assert_equal -2, -1.99.floor
  assert_equal 1.9, 1.99.floor(1)
  assert_equal -2, -1.99.floor(1)
  # 負の値も指定可能
  assert_equal 11110, 11119.floor(-1)
  assert_equal -11120, -11119.floor(-1)

  # 切り捨て（0 から 自身までの数値で、自身にもっとも近い整数または小数を返す）
  assert_equal 1, 1.99.truncate
  assert_equal -1, -1.99.truncate
  assert_equal 1.9, 1.99.truncate(1)
  assert_equal -1.9, -1.99.truncate(1)
  # 負の値も指定可能
  assert_equal 190, 199.truncate(-1)
  assert_equal -190, -199.truncate(-1)
end

# https://bugs.ruby-lang.org/issues/12512
def test_hash_transform_values
  x = {a: 1, b: 2, c: 3}
  y = x.transform_values {|v| v ** 2 }
  assert_equal({a: 1, b: 4, c: 9}, y)
  assert_equal({a: 1, b: 2, c: 3}, x)

  x.transform_values! {|v| v ** 2 }
  assert_equal({a: 1, b: 4, c: 9}, x)
end

# https://bugs.ruby-lang.org/issues/12005 --------------------
def test_unify_Fixnum_and_Bignum_into_Integer
  assert_equal Integer, 1.class
  assert_equal Integer, -1.class
  assert_equal Integer, 10000000000000000000000000000.class
  assert_equal Integer, -10000000000000000000000000000.class

  assert_equal Integer, Fixnum
  assert_equal Integer, Bignum
  a = 1
  assert a.kind_of?(Integer)
  assert a.kind_of?(Fixnum)
  assert a.kind_of?(Bignum)
  assert a.is_a?(Integer)
  assert a.is_a?(Fixnum)
  assert a.is_a?(Bignum)

  assert Integer === a
  assert Fixnum === a
  assert Bignum === a
end

class ::Fixnum
  def foo
    'foo'
  end
  def baz
    'baz'
  end
end

class ::Bignum
  def bar
    'bar'
  end
  def baz
    'baz!!'
  end
end

def test_extend_Integer
  a = 1
  assert_equal 'foo', a.foo
  assert_equal 'bar', a.bar
  assert_equal 'baz!!', a.baz
end
# https://bugs.ruby-lang.org/issues/12005 --------------------

# https://bugs.ruby-lang.org/issues/12447
def test_integer_digits
  assert_equal [3, 2, 1], 123.digits
  assert_equal [11, 7], 0x7b.digits(16)
  assert_equal [11, 7], 123.digits(16)
  assert_raises(Math::DomainError) do
    -123.digits
  end
end

# https://bugs.ruby-lang.org/issues/12300
def test_clone_with_freeze_keyword_option
  a = 'A'.freeze
  assert a.frozen?
  assert a.clone.frozen?
  refute a.clone(freeze: false).frozen?
end

# https://bugs.ruby-lang.org/issues/11999
def test_match_data_named_captures
  m = /(?<year>\d+)-(?<month>\d+)-(?<day>\d+)/.match('2016-09-01')
  assert_equal(
      {'year' => '2016', 'month' => '09', 'day' => '01'},
      m.named_captures
  )
end

# https://bugs.ruby-lang.org/issues/9179
def test_match_data_values_at
  m = /(?<year>\d+)-(?<month>\d+)-(?<day>\d+)/.match('2016-09-01')
  assert_equal(['2016', '01'], m.values_at(1, 3))
  assert_equal(['2016', '01'], m.values_at('year', 'day'))
  assert_equal(['2016', '01'], m.values_at(:year, :day))
end

# https://bugs.ruby-lang.org/issues/8110
def test_regexp_match?
  assert /\d+-\d+-\d+/.match?('2016-09-01')
  assert_nil $~
  assert '2016-09-01'.match?(/\d+-\d+-\d+/)
  assert_nil $~
  assert :hello_world.match?(/\w+_\w+/)
  assert_nil $~

  assert /\d+-\d+-\d+/.match('2016-09-01')
  assert_equal '2016-09-01', $~[0]

  assert /\d+-\d+-\d+/ =~ '2016-09-02'
  assert_equal '2016-09-02', $~[0]

  assert /\d+-\d+-\d+/ === '2016-09-03'
  assert_equal '2016-09-03', $~[0]

  assert /\w+_\w+/.match(:hello_world)
  assert_equal 'hello_world', $~[0]
end

# Regexp/String: Updated Unicode version from 8.0.0 to 9.0.0
# https://bugs.ruby-lang.org/issues/12513
def test_unicode_9_0
  assert "A\u{17000}B".match(/\p{Tangut}/)
end

# https://bugs.ruby-lang.org/issues/10085
def test_non_ascii_upcase_downcase_swapcase_capitalize
  assert_equal 'TÜRKIYE', 'Türkiye'.upcase
  assert_equal 'türkiye', 'Türkiye'.downcase
  assert_equal 'tÜRKIYE', 'Türkiye'.swapcase
  assert_equal 'Ürkiye', 'ürkiye'.capitalize

  a = 'Türkiye'
  a.upcase!
  assert_equal 'TÜRKIYE', a

  assert_equal 'TüRKIYE', 'Türkiye'.upcase(:ascii)
end

# https://bugs.ruby-lang.org/issues/12024
def test_string_with_capacity
  # https://blog.blockscore.com/new-features-in-ruby-2-4/
  append_me = ' ' * 1_000
  # String.newより2.32倍速くなるとのこと
  template  = String.new(capacity: 100_000)

  100.times { template << append_me }

  a = String.new(capacity: 1)
  a << '12345'
  assert_equal '12345', a
end

# https://bugs.ruby-lang.org/issues/11991
def test_symbol_match
  assert_equal '', :"".match(//)[0]
end

# https://bugs.ruby-lang.org/issues/6647
def test_thread_report_on_exception
  # https://blog.blockscore.com/new-features-in-ruby-2-4/
  assert_output "Starting some parallel work\nDone!\n", '' do
    puts 'Starting some parallel work'
    thread =
        Thread.new do
          sleep 0.1
          fail 'something very bad happened!'
        end
    sleep 0.2
    puts 'Done!'
  end

  assert_output "Starting some parallel work\nDone!\n", /something very bad happened!/ do
    Thread.report_on_exception = true
    puts 'Starting some parallel work'
    thread =
        Thread.new do
          sleep 0.1
          fail 'something very bad happened!'
        end
    sleep 0.2
    puts 'Done!'
  end
end

# https://bugs.ruby-lang.org/issues/11839
def test_csv_liberal_parsing_option
  require 'csv'
  input = '"Johnson, Dwayne",Dwayne "The Rock" Johnson'
  assert_raises(CSV::MalformedCSVError) { CSV.parse_line(input) }
  assert_equal(
      ['Johnson, Dwayne', 'Dwayne "The Rock" Johnson'],
      CSV.parse_line(input, liberal_parsing: true)
  )

  valid_input = '"Johnson, Dwayne","Dwayne ""The Rock"" Johnson"'
  assert_equal(
      ['Johnson, Dwayne', 'Dwayne "The Rock" Johnson'],
      CSV.parse_line(valid_input)
  )
end

# https://bugs.ruby-lang.org/issues/12224
def test_logger_options
  require 'logger'
  formatter = proc { |severity, timestamp, progname, msg| "#{severity}:#{msg}\n\n" }
  logger = Logger.new(
      STDERR,
      level: :info,
      progname: :progname,
      formatter: formatter,
      datetime_format: "%d%b%Y@%H:%M:%S"
  )
  assert_equal Logger::INFO, logger.level
  assert_equal :progname, logger.progname
  assert_equal formatter, logger.formatter
  assert_equal "%d%b%Y@%H:%M:%S", logger.datetime_format
end

# https://bugs.ruby-lang.org/issues/10772
def test_logger_shift_period_suffix
  require 'logger'
  log_path = File.expand_path('../development.log', __FILE__)
  # development.log.2016-09-01
  # development.log.2016-09-01.1
  # development.log.2016-09-01.2
  # のような名前のログファイルが出力される
  assert Logger.new(log_path, shift_period_suffix: '%Y-%m-%d')
ensure
  File.unlink(log_path) if File.exist?(log_path)
end

# OpenSSL is extracted as a gem
# https://bugs.ruby-lang.org/issues/9612
# No test.

# https://bugs.ruby-lang.org/issues/11191
def test_optparse_into_option
  # https://blog.blockscore.com/new-features-in-ruby-2-4/
  require 'optparse'
  require 'optparse/date'
  require 'optparse/uri'

  cli =
      OptionParser.new do |options|
        options.define '--from=DATE',    Date
        options.define '--url=ENDPOINT', URI
        options.define '--names=LIST',   Array
      end

  config = {}

  args = %w[
    --from  2016-02-03
    --url   https://blog.blockscore.com/
    --names John,Daniel,Delmer
  ]

  cli.parse(args, into: config)

  assert_equal(
      {
          from: Date.parse('2016-02-03'),
          url: URI('https://blog.blockscore.com/'),
          names: %w(John Daniel Delmer)
      },
      config
  )
end

# https://bugs.ruby-lang.org/issues/12189
def test_date_time_to_time
  require 'date'
  cet_date_time = DateTime.strptime('2015-11-12 CET', '%Y-%m-%d %Z')
  assert_equal '2015-11-12 00:00:00 +0100', cet_date_time.to_time.to_s
end

# https://bugs.ruby-lang.org/issues/12271
def test_time_to_time
  require 'time'
  cet_time = Time.new(2005, 2, 21, 10, 11, 12, '+01:00')
  assert_equal '2005-02-21 10:11:12 +0100', cet_time.to_time.to_s
end

# Tk is removed from stdlib.
# https://github.com/ruby/tk is the new upstream.
# https://bugs.ruby-lang.org/issues/8539
# No test.

def test_dead_lock
  begin
    Thread.current.name = "MainThread!"
    z = Thread.new{Thread.stop}
    a, b = Thread.new { 1 until b; b.join }, Thread.new { 1 until a; a.join }
    a.name = "aaaaa"
    b.name = "bbbbb"
    z.name = "zzzz"
    a.join

    flunk 'should dead lock'
  rescue Exception => e
    error_info = e.inspect
    assert /No live threads left. Deadlock\?/ =~ error_info
    assert error_info.lines.size > 50
  end
end

# preview 3 features -------------------------------
def test_rescue_modifier_syntax
  assert ('1'.to_i rescue nil)
end

# https://bugs.ruby-lang.org/issues/12548
def test_round_half
  assert_equal(13.0, 12.5.round)
  assert_equal(14.0, 13.5.round)
  assert_equal(12.0, 12.5.round(half: :even))
  assert_equal(14.0, 13.5.round(half: :even))
  assert_equal(13.0, 12.5.round(half: :up))
  assert_equal(14.0, 13.5.round(half: :up))

  assert_equal(120, 125.round(-1, half: :even))
  assert_equal(140, 135.round(-1, half: :even))

  assert_equal(12.0, (125r / 10).round(half: :even))
  assert_equal(14.0, (135r / 10).round(half: :even))
end

# https://bugs.ruby-lang.org/issues/12039
def test_finite_infinite
  assert 1.finite?
  refute 1.infinite?
  assert 1.0.finite?
  refute 1.0.infinite?
  assert 1r.finite?
  refute 1r.infinite?
  assert Complex(1).finite?
  refute Complex(1).infinite?
end

# https://bugs.ruby-lang.org/issues/12333
def test_array_concat
  assert [1, 2, 3, 4, 5, 6], [1, 2].concat([3, 4], [5, 6])
end

# https://bugs.ruby-lang.org/issues/2172
def test_numerable_chunk
  assert_instance_of Enumerator, [1, 2, 3].chunk
end

# https://bugs.ruby-lang.org/issues/11818
def test_hash_compact
  hash = { a: 1, b: nil, c: 2 }
  assert_equal({ a: 1, c: 2 }, hash.compact)
  assert_equal({ a: 1, b: nil, c: 2 }, hash)
  assert_equal({ a: 1, c: 2 }, hash.compact!)
  assert_equal({ a: 1, c: 2 }, hash)
end

# https://bugs.ruby-lang.org/issues/12210
def test_set_compare_by_identity
  require 'set'
  set = Set.new
  values = ['a', 'a', 'b', 'b']
  set.merge(values)
  refute set.compare_by_identity?
  assert_equal ['a', 'b'], set.to_a

  set = Set.new
  values = ['a', 'a', 'b', 'b']
  set.compare_by_identity
  set.merge(values)
  assert set.compare_by_identity?
  assert_equal ['a', 'a', 'b', 'b'], set.to_a
end

# https://bugs.ruby-lang.org/issues/12333
def test_string_concat
  s = 'a'
  assert_equal 'abc', s.concat('b', 'c')
  assert_equal 'abc', s

  s = 'z'
  assert_equal 'xyz', s.prepend('x', 'y')
  assert_equal 'xyz', s
end

# https://bugs.ruby-lang.org/issues/12596
def test_pathname_empty?
  require 'tempfile'
  require 'pathname'
  Dir.mktmpdir do |dir|
    assert Pathname(dir).empty?
  end
  Tempfile.create("foo") do |f|
    assert Pathname(f).empty?
  end
end

# https://bugs.ruby-lang.org/issues/12553
def test_io_readlines
  require 'tempfile'
  tf = Tempfile.new("foo")
  tf.print "abc\ndef\nghi\n"
  tf.close

  File.open(tf.path) do |f|
    assert_equal ["abc\n", "def\n", "ghi\n"], f.readlines
    f.rewind
    assert_equal ["abc", "def", "ghi"], f.readlines(chomp: true)
  end
ensure
  tf&.unlink
end

# https://bugs.ruby-lang.org/issues/10055
def test_shellwords_shellsplit
  require 'shellwords'
  assert_equal ['printf', '%s\n'], Shellwords.shellsplit(%q{printf "%s\\n"})
  assert_equal ['printf', '%s\n'], Shellwords.shellsplit(%q{printf '%s\\n'})
end

# https://bugs.ruby-lang.org/issues/12799
def test_ipaddr_compare
  require 'ipaddr'
  refute IPAddr.new("1.1.1.1") == "sometext"
end

def test_casecmp?
  assert 'abc'.casecmp?('ABC')
  refute 'abc'.casecmp?('DEF')
  assert :abc.casecmp?(:ABC)
  refute :abc.casecmp?(:DEF)

  assert_equal 0, 'abc'.casecmp('ABC')
  assert_equal -1, 'abc'.casecmp('DEF')
  assert_equal 0, :abc.casecmp(:ABC)
  assert_equal -1, :abc.casecmp(:DEF)
end

def test_unpack1
  assert_equal [65, 66, 67], "ABC".unpack("C*")
  assert_equal 65, "ABC".unpack1("C*")
end
end