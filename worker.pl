#!/usr/bin/env perl
use strict;
use warnings;
use Getopt::Long;
use FindBin qw($RealBin);
use lib "$RealBin/.";
use Quotes;

# Command line options
my $anime;
my $character;
GetOptions(
    "anime|a=s" => \$anime,
    "character|c=s" => \$character
);

# Filter quotes based on arguments
my @filtered_quotes = @AnimeQuotes::Quotes;

if ($anime) {
    @filtered_quotes = grep {
        $_->{anime} =~ /$anime/i
    } @filtered_quotes;
}

if ($character) {
    @filtered_quotes = grep {
        $_->{character} =~ /$character/i
    } @filtered_quotes;
}

if (@filtered_quotes == 0) {
    die "No quotes found matching your criteria.\n";
}

# Select random quote
my $quote = $filtered_quotes[rand @filtered_quotes];

# Print formatted quote
print "\"$quote->{quote}\"\n";
print "  - $quote->{character} ($quote->{anime})\n";
