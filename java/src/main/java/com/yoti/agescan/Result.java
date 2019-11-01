package com.yoti.agescan;

import lombok.Data;
import lombok.Getter;
import lombok.Setter;

@Data
public class Result {
    private double  pred_age;
    private double uncertainty;

}
