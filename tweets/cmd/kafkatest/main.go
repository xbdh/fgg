package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main(){
	addrs:=[]string{"127.0.0.1:9092"}
	consumer :=NewKafkaConsumer(addrs)

	partitionlist,err:=consumer.Partitions("tweet")
	if err!=nil{
		fmt.Println("err")
		return
	}

	for part:=range partitionlist{
		pc ,err:=consumer.ConsumePartition("tweet",int32(part),sarama.OffsetNewest)
		if err!=nil{
			fmt.Println("err")
			return
		}
		defer pc.AsyncClose()

		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg:=range pc.Messages(){
				fmt.Println(msg.Value)
			}
		}(pc)
	}



}
func NewKafkaConsumer(addrs []string) sarama.Consumer {
	c := sarama.NewConfig()
	p, err := sarama.NewConsumer(addrs, c)
	if err != nil {
		panic(err)
	}
	return p
}