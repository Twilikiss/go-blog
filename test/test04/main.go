package main

import (
	"fmt"
	"regexp"
	"strings"
)

// TrimHtml 去除字符串中的html标签
func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "")
	return strings.TrimSpace(src)
}

func main() {
	data := "<h2 id=\"h2-1-redis-\"><a name=\"1. Redis是什么？简述它的优缺点？\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>1. Redis是什么？简述它的优缺点？</h2><p>Redis本质上是一个<strong>Key-Value类型的内存数据库</strong>，很像Memcached，整个数据库加载在内存当中操作，定期通过异步操作把数据库中的数据flush到硬盘上进行保存。</p>\n<p><strong>因为是纯内存操作</strong>，Redis的性能非常出色，每秒可以处理超过 10万次读写操作，是已知性能最快的Key-Value 数据库。</p>\n<p><strong>优点</strong>：</p>\n<ul>\n<li>读写性能极高， Redis能读的速度是110000次/s，写的速度是81000次/s。</li><li>支持数据持久化，支持AOF和RDB两种持久化方式。</li><li>支持事务， Redis的所有操作都是原子性的，意思就是要么成功执行要么失败完全不执行。单个操作是原子性的。多个操作也支持事务，即原子性，通过MULTI和EXEC指令包起来。</li><li>数据结构丰富，除了支持string类型的value外，还支持hash、set、zset、list等数据结构。</li><li>支持主从复制，主机会自动将数据同步到从机，可以进行读写分离。</li><li>丰富的特性 – Redis还支持 publish/subscribe， 通知， key 过期等特性。</li></ul>\n<p><strong>缺点</strong>：</p>\n<ul>\n<li>数据库容量受到<strong>物理内存的限制</strong>，不能用作海量数据的高性能读写，因此Redis适合的场景主要局限在较小数据量的高性能操作和运算上。</li><li><strong>主机宕机，宕机前有部分数据未能及时同步到从机</strong>，切换IP后还会引入数据不一致的问题，降低了系统的可用性。</li></ul>\n<h2 id=\"h2-2-redis-\"><a name=\"2. Redis为什么这么快？\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>2. Redis为什么这么快？</h2><ul>\n<li>内存存储：Redis是使用内存(in-memeroy)存储，没有磁盘IO上的开销。数据存在内存中，类似于 HashMap，HashMap 的优势就是查找和操作的时间复杂度都是O(1)。</li><li>单线程实现（ Redis 6.0以前）：Redis使用单个线程处理请求，避免了多个线程之间线程切换和锁资源争用的开销。注意：单线程是指的是在核心网络模型中，网络请求模块使用一个线程来处理，即一个线程处理所有网络请求。</li><li>非阻塞IO：Redis使用多路复用IO技术，将epoll作为I/O多路复用技术的实现，再加上Redis自身的事件处理模型将epoll中的连接、读写、关闭都转换为事件，不在网络I/O上浪费过多的时间。</li><li>优化的数据结构：Redis有诸多可以直接应用的优化数据结构的实现，应用层可以直接使用原生的数据结构提升性能。</li><li>使用底层模型不同：Redis直接自己构建了 VM (虚拟内存)机制 ，因为一般的系统调用系统函数的话，会浪费一定的时间去移动和请求。Redis的VM(虚拟内存)机制就是暂时把不经常访问的数据(冷数据)从内存交换到磁盘中，从而腾出宝贵的内存空间用于其它需要访问的数据(热数据)。通过VM功能可以实现冷热数据分离，使热数据仍在内存中、冷数据保存到磁盘。这样就可以避免因为内存不足而造成访问速度下降的问题。Redis提高数据库容量的办法有两种：一种是可以将数据分割到多个RedisServer上；另一种是使用虚拟内存把那些不经常访问的数据交换到磁盘上。<strong>需要特别注意的是Redis并没有使用OS提供的Swap，而是自己实现。</strong></li></ul>\n<h2 id=\"h2-3-redis-memcached-\"><a name=\"3. Redis相比Memcached有哪些优势？\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>3. Redis相比Memcached有哪些优势？</h2><ul>\n<li>数据类型：Memcached所有的值均是简单的字符串，Redis支持更为丰富的数据类型，支持string(字符串)，list(列表)，Set(集合)、Sorted Set(有序集合)、Hash(哈希)等。</li><li>持久化：Redis支持数据落地持久化存储，可以将内存中的数据保持在磁盘中，重启的时候可以再次加载进行使用。 memcache不支持数据持久存储 。</li><li>集群模式：Redis提供主从同步机制，以及 Cluster集群部署能力，能够提供高可用服务。Memcached没有原生的集群模式，需要依靠客户端来实现往集群中分片写入数据</li><li>性能对比：Redis的速度比Memcached快很多。</li><li>网络IO模型：Redis使用单线程的多路 IO 复用模型，Memcached使用多线程的非阻塞IO模式。</li><li>Redis支持服务器端的数据操作：Redis相比Memcached来说，拥有更多的数据结构和并支持更丰富的数据操作，通常在Memcached里，你需要将数据拿到客户端来进行类似的修改再set回去。这大大增加了网络IO的次数和数据体积。在Redis中，这些复杂的操作通常和一般的GET/SET一样高效。所以，如果需要缓存能够支持更复杂的结构和操作，那么Redis会是不错的选择。</li></ul>\n<h2 id=\"h2-4-redis-\"><a name=\"4. 为什么要用 Redis 做缓存？\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>4. 为什么要用 Redis 做缓存？</h2><p><strong>从高并发上来说：</strong></p>\n<ul>\n<li>直接操作缓存能够承受的请求是远远大于直接访问数据库的，所以我们可以考虑把数据库中的部分数据转移到缓存中去，这样用户的一部分请求会直接到缓存这里而不用经过数据库。</li></ul>\n<p><strong>从高性能上来说：</strong></p>\n<ul>\n<li>用户第一次访问数据库中的某些数据。 因为是从硬盘上读取的所以这个过程会比较慢。将该用户访问的数据存在缓存中，下一次再访问这些数据的时候就可以直接从缓存中获取了。操作缓存就是直接操作内存，所以速度相当快。如果数据库中的对应数据改变的之后，同步改变缓存中相应的数据。</li></ul>\n<h2 id=\"h2-5-redis-map-guava-\"><a name=\"5. 为什么要用 Redis 而不用 map/guava 做缓存?\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>5. 为什么要用 Redis 而不用 map/guava 做缓存?</h2><p>缓存分为本地缓存和分布式缓存。以java为例，使用自带的map或者guava实现的是本地缓存，最主要的特点是轻量以及快速，生命周期随着jvm的销毁而结束，并且在多实例的情况下，每个实例都需要各自保存一份缓存，缓存不具有一致性。</p>\n<p>使用Redis或memcached之类的称为分布式缓存，在多实例的情况下，各实例共用一份缓存数据，缓存具有一致性。缺点是需要保持Redis或memcached服务的高可用，整个程序架构上较为复杂。</p>\n<p>对比:</p>\n<ul>\n<li>Redis 可以用几十 G 内存来做缓存，Map 不行，一般 JVM 也就分几个 G 数据就够大了；</li><li>Redis 的缓存可以持久化，Map 是内存对象，程序一重启数据就没了；</li><li>Redis 可以实现分布式的缓存，Map 只能存在创建它的程序里；</li><li>Redis 可以处理每秒百万级的并发，是专业的缓存服务，Map 只是一个普通的对象；</li><li>Redis 缓存有过期机制，Map 本身无此功能；Redis 有丰富的 API，Map 就简单太多了；</li><li>Redis可单独部署，多个项目之间可以共享，本地内存无法共享；</li><li>Redis有专门的管理工具可以查看缓存数据。</li></ul>"
	fmt.Println(string([]rune(TrimHtml(data))[:150]))
}
